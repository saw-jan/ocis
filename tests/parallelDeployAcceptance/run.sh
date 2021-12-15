#!/usr/bin/env bash
[[ "${DEBUG}" == "true" ]] && set -x

SCRIPT_PATH="`dirname \"$0\"`"
SCRIPT_PATH="`( cd \"${SCRIPT_PATH}\" && pwd )`"

OCIS_PATH="${SCRIPT_PATH}/../../"

BEHAT_YML="${SCRIPT_PATH}/config/behat.yml"

# use bootstrap from core repo
# export BEHAT_PARAMS='{"autoload": "'${PATH_TO_CORE}'/tests/acceptance/features/bootstrap"}'

# Allow optionally passing in the path to the behat program.
if [ -z "${BEHAT_BIN}" ]
then
    BEHAT=${OCIS_PATH}vendor-bin/behat/vendor/bin/behat
else
    BEHAT=${BEHAT_BIN}
fi

# Command line options processed here will override environment variables that
# might have been set by the caller, or in the code above.
while [[ $# -gt 0 ]]
do
	key="$1"
	case ${key} in
		--type)
			# Lowercase the parameter value, so the user can provide "API", "CLI", "webUI" etc
			ACCEPTANCE_TEST_TYPE="${2,,}"
			shift
			;;
		*)
			# A "random" parameter is presumed to be a feature file to run.
			# Typically that will be specified at the end, or as the only
			# parameter.
			BEHAT_FEATURE="$1"
			;;
	esac
	shift
done

if [[ "${BEHAT_SUITE}" == api* ]] || [ "${ACCEPTANCE_TEST_TYPE}" = "api" ]
then
	TEST_TYPE_TAG="@api"
	TEST_TYPE_TEXT="API"
	RUNNING_API_TESTS=true
fi


BEHAT_SUITE_OPTION=""

if [ -z "$BEHAT_FILTER_TAGS" ];
then
    BEHAT_FILTER_TAGS=""
fi

function run_behat_tests() {
	echo "Running '${SUITE_FEATURE_TEXT}' tests tagged '${BEHAT_FILTER_TAGS}':"

	${BEHAT} --colors -c ${BEHAT_YML} -f junit -f pretty ${BEHAT_SUITE_OPTION} --tags ${BEHAT_FILTER_TAGS} ${BEHAT_FEATURE}
}

if [ -z "${BEHAT_SUITE}" ] && [ -z "${BEHAT_FEATURE}" ]
then
	SUITE_FEATURE_TEXT="all ${TEST_TYPE_TEXT}"
	run_behat_tests
else
	if [ -n "${BEHAT_SUITE}" ]
	then
		SUITE_FEATURE_TEXT="${BEHAT_SUITE}"
		BEHAT_SUITE_OPTION="--suite=${SUITE_FEATURE_TEXT}"
	fi

	if [ -n "${BEHAT_FEATURE}" ]
	then
		# If running a whole feature, it will be something like login.feature
		# If running just a single scenario, it will also have the line number
		# like login.feature:36 - which will be parsed correctly like a "file"
		# by basename.
		BEHAT_FEATURE_FILE=`basename ${BEHAT_FEATURE}`
		SUITE_FEATURE_TEXT="${SUITE_FEATURE_TEXT}/${BEHAT_FEATURE_FILE}"
	fi
	run_behat_tests
fi