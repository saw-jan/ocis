<?php declare(strict_types=1);
/**
 * ownCloud
 *
 * @author Kiran Parajuli <kiran@jankaritech.com>
 * @copyright Copyright (c) 2022 Kiran Parajuli kiran@jankaritech.com
 */

namespace TestHelpers;

use GuzzleHttp\Exception\GuzzleException;
use Psr\Http\Message\RequestInterface;
use Psr\Http\Message\ResponseInterface;

/**
 * A helper class for managing Graph API requests
 */
class GraphHelper {
	/**
	 * @return string[]
	 */
	private static function getRequestHeaders(): array {
		return [
			'Content-Type' => 'application/json',
		];
	}

	/**
	 *
	 * @return string
	 */
	public static function getUUIDv4Regex(): string {
		return '[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}';
	}

	/**
	 * @param string $id
	 *
	 * @return bool
	 */
	public static function isUUIDv4(string $id): bool {
		$regex = "/^" . self::getUUIDv4Regex() . "$/i";
		return (bool)preg_match($regex, $id);
	}

	/**
	 * @param string $spaceId
	 *
	 * @return bool
	 */
	public static function isSpaceId(string $spaceId): bool {
		$regex = "/^" . self::getUUIDv4Regex() . '\\$' . self::getUUIDv4Regex() . "$/i";
		return (bool)preg_match($regex, $spaceId);
	}

	/**
	 * @return string
	 */
	public static function getSpaceIdRegex(): string {
		return self::getUUIDv4Regex() . '\\\$' . self::getUUIDv4Regex();
	}

	/**
	 * @return string
	 */
	public static function getShareIdRegex(): string {
		return self::getUUIDv4Regex() . ':' . self::getUUIDv4Regex() . ':' . self::getUUIDv4Regex();
	}

	/**
	 * Key name can consist of @@@
	 * This function separate such key and return its actual value from actual drive response which can be used for assertion
	 *
	 * @param string $keyName
	 * @param array $actualDriveInformation
	 *
	 * @return string
	 */
	public static function separateAndGetValueForKey(string $keyName, array $actualDriveInformation): string {
		// break the segment with @@@ to find the actual value from the actual drive information
		$separatedKey = explode("@@@", $keyName);
		// this stores the actual value of each key from drive information response used for assertion
		$actualKeyValue = $actualDriveInformation;

		foreach ($separatedKey as $key) {
			$actualKeyValue = $actualKeyValue[$key];
		}

		return $actualKeyValue;
	}

	/**
	 * @param string $baseUrl
	 * @param string $path
	 *
	 * @return string
	 */
	public static function getFullUrl(string $baseUrl, string $path): string {
		$fullUrl = $baseUrl;
		if (\substr($fullUrl, -1) !== '/') {
			$fullUrl .= '/';
		}
		$fullUrl .= 'graph/v1.0/' . $path;
		return $fullUrl;
	}

	/**
	 * @param string $baseUrl
	 * @param string $path
	 *
	 * @return string
	 */
	public static function getBetaFullUrl(string $baseUrl, string $path): string {
		$baseUrl = rtrim($baseUrl, "/");
		return $baseUrl . '/graph/v1beta1/' . $path;
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $method
	 * @param string $path
	 * @param string|null $body
	 * @param array|null $headers
	 *
	 * @return RequestInterface
	 */
	public static function createRequest(
		string $baseUrl,
		string $xRequestId,
		string $method,
		string $path,
		?string $body = null,
		?array $headers = []
	): RequestInterface {
		$fullUrl = self::getFullUrl($baseUrl, $path);
		return HttpRequestHelper::createRequest(
			$fullUrl,
			$xRequestId,
			$method,
			$headers,
			$body
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userName
	 * @param string $password
	 * @param string|null $email
	 * @param string|null $displayName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function createUser(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userName,
		string $password,
		?string $email = null,
		?string $displayName = null
	): ResponseInterface {
		$payload = self::prepareCreateUserPayload(
			$userName,
			$password,
			$email,
			$displayName
		);

		$url = self::getFullUrl($baseUrl, 'users');
		return HttpRequestHelper::post(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
			$payload
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userId
	 * @param string $method
	 * @param string|null $userName
	 * @param string|null $password
	 * @param string|null $email
	 * @param string|null $displayName
	 * @param bool|true $accountEnabled
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function editUser(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userId,
		?string $method = "PATCH",
		?string $userName = null,
		?string $password = null,
		?string $email = null,
		?string $displayName = null,
		?bool $accountEnabled = true
	): ResponseInterface {
		$payload = self::preparePatchUserPayload(
			$userName,
			$password,
			$email,
			$displayName,
			$accountEnabled
		);
		$url = self::getFullUrl($baseUrl, 'users/' . $userId);
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			$method,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
			$payload
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUser(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userName);
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $userPassword
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getOwnInformationAndGroupMemberships(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $userPassword
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'me/?%24expand=memberOf');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$userPassword,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function deleteUser(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userName);
		return HttpRequestHelper::delete(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function deleteUserByUserId(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userId);

		return HttpRequestHelper::delete(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $byUser
	 * @param string $userPassword
	 * @param string|null $user
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUserWithDriveInformation(
		string $baseUrl,
		string $xRequestId,
		string $byUser,
		string $userPassword,
		?string $user = null
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $user . '?%24select=&%24expand=drive');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$byUser,
			$userPassword,
		);
	}

	/***
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $byUser
	 * @param string $userPassword
	 * @param string $userId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getPersonalDriveInformationByUserId(
		string $baseUrl,
		string $xRequestId,
		string $byUser,
		string $userPassword,
		string $userId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userId . '/drive');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$byUser,
			$userPassword
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $byUser
	 * @param string $userPassword
	 * @param string|null $user
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUserWithGroupInformation(
		string $baseUrl,
		string $xRequestId,
		string $byUser,
		string $userPassword,
		?string $user = null
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $user . '?%24expand=memberOf');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$byUser,
			$userPassword,
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $groupName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function createGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $groupName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups');
		$payload['displayName'] = $groupName;
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"POST",
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $groupId
	 * @param string $displayName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function updateGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $groupId,
		string $displayName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupId);
		$payload['displayName'] = $displayName;
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"PATCH",
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUsers(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getGroups(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $groupName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $groupName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupName);
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function deleteGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $groupId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupId);
		return HttpRequestHelper::delete(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
		);
	}

	/**
	 * add multiple users to a group at once
	 *
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $groupId
	 * @param array $userIds
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function addUsersToGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $groupId,
		array $userIds
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupId);
		$payload = [ "members@odata.bind" => [] ];
		foreach ($userIds as $userId) {
			$payload["members@odata.bind"][] = self::getFullUrl($baseUrl, 'users/' . $userId);
		}
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			'PATCH',
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userId
	 * @param string $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function addUserToGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userId,
		string $groupId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupId . '/members/$ref');
		$body = [
			"@odata.id" => self::getFullUrl($baseUrl, 'users/' . $userId)
		];
		return HttpRequestHelper::post(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
			self::getRequestHeaders(),
			\json_encode($body)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $userId
	 * @param string $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function removeUserFromGroup(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $userId,
		string $groupId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupId . '/members/' . $userId . '/$ref');
		return HttpRequestHelper::delete(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword,
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getMembersList(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		string $groupId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'groups/' . $groupId . '/members');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword
		);
	}

	/**
	 * returns single group information along with its member information when groupId is provided
	 * else return all group information along with its member information
	 *
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $adminUser
	 * @param string $adminPassword
	 * @param string|null $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getSingleOrAllGroupsAlongWithMembers(
		string $baseUrl,
		string $xRequestId,
		string $adminUser,
		string $adminPassword,
		?string $groupId = null
	): ResponseInterface {
		// we can expand to get list of members for a single group with groupId and also expand to get all groups with all its members
		$endPath = ($groupId) ? '/' . $groupId . '?$expand=members' : '?$expand=members';
		$url = self::getFullUrl($baseUrl, 'groups' . $endPath);
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$adminUser,
			$adminPassword
		);
	}

	/**
	 * returns json encoded payload for user creating request
	 *
	 * @param string|null $userName
	 * @param string|null $password
	 * @param string|null $email
	 * @param string|null $displayName
	 *
	 * @return string
	 */
	public static function prepareCreateUserPayload(
		string $userName,
		string $password,
		?string $email,
		?string $displayName
	): string {
		$payload['onPremisesSamAccountName'] = $userName;
		$payload['passwordProfile'] = ['password' => $password];
		$payload['displayName'] = $displayName ?? $userName;
		if (empty($email)) {
			$payload['mail'] = $userName . '@example.com';
		} else {
			$payload['mail'] = $email;
		}
		$payload['accountEnabled'] = true;
		return \json_encode($payload);
	}

	/**
	 * returns encoded json payload for user patching requests
	 *
	 * @param string|null $userName
	 * @param string|null $password
	 * @param string|null $email
	 * @param string|null $displayName
	 * @param bool|true $accountEnabled
	 *
	 * @return string
	 */
	public static function preparePatchUserPayload(
		?string $userName,
		?string $password,
		?string $email,
		?string $displayName,
		?bool $accountEnabled
	): string {
		$payload = [];
		if ($userName !== null) {
			// comment on after fixing #5755 because now it crashes server
			// if (empty($userName)) {
			//   $payload['onPremisesSamAccountName'] = ' ';
			// } else $payload['onPremisesSamAccountName'] = $userName;
			$payload['onPremisesSamAccountName'] = $userName;
		}
		if ($password !== null) {
			$payload['passwordProfile'] = ['password' => $password];
		}
		if ($displayName !== null) {
			$payload['displayName'] = $displayName;
		}
		if ($email !== null) {
			$payload['mail'] = $email;
		}
		$payload['accountEnabled'] = $accountEnabled;
		return \json_encode($payload);
	}

	/**
	 * Send Graph Create Space Request
	 *
	 * @param string $baseUrl
	 * @param  string $user
	 * @param  string $password
	 * @param  string $body
	 * @param  string $xRequestId
	 * @param  array  $headers
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function createSpace(
		string $baseUrl,
		string $user,
		string $password,
		string $body,
		string $xRequestId = '',
		array $headers = []
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives');

		return HttpRequestHelper::post($url, $xRequestId, $user, $password, $headers, $body);
	}

	/**
	 * Send Graph Update Space Request
	 *
	 * @param  string $baseUrl
	 * @param  string $user
	 * @param  string $password
	 * @param  mixed $body
	 * @param  string $spaceId
	 * @param  string $xRequestId
	 * @param  array  $headers
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function updateSpace(
		string $baseUrl,
		string $user,
		string $password,
		$body,
		string $spaceId,
		string $xRequestId = '',
		array $headers = []
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives/' . $spaceId);

		return HttpRequestHelper::sendRequest($url, $xRequestId, 'PATCH', $user, $password, $headers, $body);
	}

	/**
	 * Send Graph List My Spaces Request
	 *
	 * @param  string $baseUrl
	 * @param  string $user
	 * @param  string $password
	 * @param  string $urlArguments
	 * @param  string $xRequestId
	 * @param  array  $body
	 * @param  array  $headers
	 *
	 * @return ResponseInterface
	 *
	 * @throws GuzzleException
	 */
	public static function getMySpaces(
		string $baseUrl,
		string $user,
		string $password,
		string $urlArguments = '',
		string $xRequestId = '',
		array  $body = [],
		array  $headers = []
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'me/drives/' . $urlArguments);

		return HttpRequestHelper::get($url, $xRequestId, $user, $password, $headers, $body);
	}

	/**
	 * Send Graph List All Spaces Request
	 *
	 * @param  string $baseUrl
	 * @param  string $user
	 * @param  string $password
	 * @param  string $urlArguments
	 * @param  string $xRequestId
	 * @param  array  $body
	 * @param  array  $headers
	 *
	 * @return ResponseInterface
	 *
	 * @throws GuzzleException
	 */
	public static function getAllSpaces(
		string $baseUrl,
		string $user,
		string $password,
		string $urlArguments = '',
		string $xRequestId = '',
		array  $body = [],
		array  $headers = []
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives/' . $urlArguments);

		return HttpRequestHelper::get($url, $xRequestId, $user, $password, $headers, $body);
	}

	/**
	 * Send Graph List Single Space Request
	 *
	 * @param string $baseUrl
	 * @param string $user
	 * @param string $password
	 * @param string $spaceId
	 * @param string $urlArguments
	 * @param string $xRequestId
	 * @param array $body
	 * @param array $headers
	 *
	 * @return ResponseInterface
	 *
	 * @throws GuzzleException
	 */
	public static function getSingleSpace(
		string $baseUrl,
		string $user,
		string $password,
		string $spaceId,
		string $urlArguments = '',
		string $xRequestId = '',
		array  $body = [],
		array  $headers = []
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives/' . $spaceId . "/" . $urlArguments);

		return HttpRequestHelper::get($url, $xRequestId, $user, $password, $headers, $body);
	}

	/**
	 * send disable space request
	 *
	 * @param string $baseUrl
	 * @param string $user
	 * @param string $password
	 * @param string $spaceId
	 * @param string $xRequestId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function disableSpace(
		string $baseUrl,
		string $user,
		string $password,
		string $spaceId,
		string $xRequestId = ''
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives/' . $spaceId);

		return HttpRequestHelper::delete(
			$url,
			$xRequestId,
			$user,
			$password
		);
	}

	/**
	 * send delete space request
	 *
	 * @param string $baseUrl
	 * @param string $user
	 * @param string $password
	 * @param string $spaceId
	 * @param string $xRequestId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function deleteSpace(
		string $baseUrl,
		string $user,
		string $password,
		string $spaceId,
		string $xRequestId = ''
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives/' . $spaceId);
		$header = ["Purge" => "T"];

		return HttpRequestHelper::delete(
			$url,
			$xRequestId,
			$user,
			$password,
			$header
		);
	}

	/**
	 * Send restore Space Request
	 *
	 * @param  string $baseUrl
	 * @param  string $user
	 * @param  string $password
	 * @param  string $spaceId
	 * @param string $xRequestId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function restoreSpace(
		string $baseUrl,
		string $user,
		string $password,
		string $spaceId,
		string $xRequestId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'drives/' . $spaceId);
		$header = ["restore" => true];
		$body = '{}';

		return HttpRequestHelper::sendRequest($url, $xRequestId, 'PATCH', $user, $password, $header, $body);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $currentPassword
	 * @param string $newPassword
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function changeOwnPassword(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $currentPassword,
		string $newPassword
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'me/changePassword');
		$payload['currentPassword'] = $currentPassword;
		$payload['newPassword'] = $newPassword;

		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"POST",
			$user,
			$password,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $user
	 * @param string $password
	 * @param string $xRequestId
	 * @param array $body
	 * @param array $headers
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getTags(
		string $baseUrl,
		string $user,
		string $password,
		string $xRequestId = '',
		array  $body = [],
		array  $headers = []
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'extensions/org.libregraph/tags');

		return HttpRequestHelper::get($url, $xRequestId, $user, $password, $headers, $body);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $resourceId
	 * @param array $tagName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function createTags(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $resourceId,
		array $tagName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'extensions/org.libregraph/tags');
		$payload['resourceId'] = $resourceId;
		$payload['tags'] = $tagName;

		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"PUT",
			$user,
			$password,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $resourceId
	 * @param array $tagName
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function deleteTags(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $resourceId,
		array $tagName
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'extensions/org.libregraph/tags');
		$payload['resourceId'] = $resourceId;
		$payload['tags'] = $tagName;

		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"DELETE",
			$user,
			$password,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getApplications(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'applications');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUsersWithFilterMemberOf(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $groupId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users' . '?$filter=memberOf/any(m:m/id ' . "eq '$groupId')");
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param array $groupIdArray
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUsersOfTwoGroups(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		array $groupIdArray
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users' . '?$filter=memberOf/any(m:m/id ' . "eq '$groupIdArray[0]') " . "and memberOf/any(m:m/id eq '$groupIdArray[1]')");
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $firstGroup
	 * @param string $secondGroup
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUsersFromOneOrOtherGroup(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $firstGroup,
		string $secondGroup
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users' . '?$filter=memberOf/any(m:m/id ' . "eq '$firstGroup') " . "or memberOf/any(m:m/id eq '$secondGroup')");
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $roleId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUsersWithFilterRoleAssignment(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $roleId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users' . '?$filter=appRoleAssignments/any(m:m/appRoleId ' . "eq '$roleId')");
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $roleId
	 * @param string $groupId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getUsersWithFilterRolesAssignmentAndMemberOf(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $roleId,
		string $groupId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users' . '?$filter=appRoleAssignments/any(m:m/appRoleId ' . "eq '$roleId') " . "and memberOf/any(m:m/id eq '$groupId')");
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $appRoleId
	 * @param string $applicationId
	 * @param string $userId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function assignRole(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $appRoleId,
		string $applicationId,
		string $userId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userId . '/appRoleAssignments');
		$payload['principalId'] = $userId;
		$payload['appRoleId'] = $appRoleId;
		$payload['resourceId'] = $applicationId;
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"POST",
			$user,
			$password,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $userId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getAssignedRole(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $userId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userId . '/appRoleAssignments');
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $userId
	 * @param string $path
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function generateGDPRReport(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $userId,
		string $path
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userId . '/exportPersonalData');
		// this payload is the storage location of the report generated
		$payload['storageLocation'] = $path;
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"POST",
			$user,
			$password,
			self::getRequestHeaders(),
			\json_encode($payload)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $appRoleAssignmentId
	 * @param string $userId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function unassignRole(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $appRoleAssignmentId,
		string $userId
	): ResponseInterface {
		$url = self::getFullUrl($baseUrl, 'users/' . $userId . '/appRoleAssignments/' . $appRoleAssignmentId);
		return HttpRequestHelper::sendRequest(
			$url,
			$xRequestId,
			"DELETE",
			$user,
			$password,
			self::getRequestHeaders(),
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $path
	 *
	 * @return string
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function getShareMountId(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $path
	): string {
		$response = self::getMySpaces(
			$baseUrl,
			$user,
			$password,
			'',
			$xRequestId
		);
		$drives = json_decode($response->getBody()->getContents(), true, 512, JSON_THROW_ON_ERROR);

		// the response returns the shared resource in driveAlias all in lowercase,
		// For example: if we get the property of a shared resource "FOLDER" then the response contains "driveAlias": "mountpoint/folder"
		// In case of two shares with same name, the response for the second shared resource will contain, "driveAlias": "mountpoint/folder-(2)"
		$path = strtolower($path);
		foreach ($drives["value"] as $value) {
			if ($value["driveAlias"] === "mountpoint/" . $path) {
				return $value["id"];
			}
		}
		throw new \Exception(
			__METHOD__
			. " Cannot find share mountpoint id of '$path' for user '$user'"
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $language
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function switchSystemLanguage(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $language
	): ResponseInterface {
		$fullUrl = self::getFullUrl($baseUrl, 'me');
		$payload['preferredLanguage'] = $language;
		return HttpRequestHelper::sendRequest($fullUrl, $xRequestId, 'PATCH', $user, $password, null, \json_encode($payload));
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $spaceId
	 * @param string $itemId
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function getPermissionsList(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $spaceId,
		string $itemId
	): ResponseInterface {
		$url = self::getBetaFullUrl($baseUrl, "drives/$spaceId/items/$itemId/permissions");
		return HttpRequestHelper::get(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders()
		);
	}

	/**
	 * Get the role id by name
	 *
	 * @param string $role
	 *
	 * @return string
	 *
	 */
	public static function getRoleIdByName(
		string $role
	): string {
		switch ($role) {
			case 'Viewer':
				return 'b1e2218d-eef8-4d4c-b82d-0f1a1b48f3b5';
			case 'Space Viewer':
				return 'a8d5fe5e-96e3-418d-825b-534dbdf22b99';
			case 'Editor':
				return 'fb6c3e19-e378-47e5-b277-9732f9de6e21';
			case 'Space Editor':
				return '58c63c02-1d89-4572-916a-870abc5a1b7d';
			case 'File Editor':
				return '2d00ce52-1fc2-4dbc-8b95-a73b73395f5a';
			case 'Co Owner':
				return '3a4ba8e9-6a0d-4235-9140-0e7a34007abe';
			case 'Uploader':
				return '1c996275-f1c9-4e71-abdf-a42f6495e960';
			case 'Manager':
				return '312c0871-5ef7-4b3a-85b6-0e4074c64049';
		}
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $spaceId
	 * @param string $itemId
	 * @param string $shareeId
	 * @param string $shareType
	 * @param string|null $role
	 *
	 * @return ResponseInterface
	 * @throws \JsonException
	 */
	public static function sendSharingInvitation(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $spaceId,
		string $itemId,
		string $shareeId,
		string $shareType,
		?string $role
	): ResponseInterface {
		$url = self::getBetaFullUrl($baseUrl, "drives/$spaceId/items/$itemId/invite");
		$body = [];

		$recipients['objectId'] = $shareeId;
		$recipients['@libre.graph.recipient.type'] = $shareType;

		$body['recipients'] = [$recipients];

		if ($role !== null) {
			$roleId = self::getRoleIdByName($role);
			$body['roles'] = [$roleId];
		}

		return HttpRequestHelper::post(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders(),
			\json_encode($body)
		);
	}

	/**
	 * @param string $baseUrl
	 * @param string $xRequestId
	 * @param string $user
	 * @param string $password
	 * @param string $spaceId
	 * @param string $itemId
	 * @param mixed $body
	 *
	 * @return ResponseInterface
	 * @throws GuzzleException
	 */
	public static function createLinkShare(
		string $baseUrl,
		string $xRequestId,
		string $user,
		string $password,
		string $spaceId,
		string $itemId,
		$body
	): ResponseInterface {
		$url = self::getBetaFullUrl($baseUrl, "drives/$spaceId/items/$itemId/createLink");
		return HttpRequestHelper::post(
			$url,
			$xRequestId,
			$user,
			$password,
			self::getRequestHeaders(),
			$body
		);
	}
}
