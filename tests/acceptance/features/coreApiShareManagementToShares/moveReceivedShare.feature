@skipOnReva @issue-1289 @issue-1328
Feature: sharing

  Background:
    Given using OCS API version "1"
    And these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
      | Carol    |


  Scenario: keep user/group shares when the user renames the share within the Shares folder
    Given group "grp1" has been created
    And user "Brian" has been added to group "grp1"
    And user "Carol" has been added to group "grp1"
    And user "Alice" has created folder "/TMP"
    When user "Alice" shares folder "TMP" with group "grp1" using the sharing API
    And user "Brian" moves folder "/Shares/TMP" to "/Shares/new" using the WebDAV API
    And the administrator deletes user "Carol" using the provisioning API
    Then the OCS status code of responses on all endpoints should be "100"
    And the HTTP status code of responses on each endpoint should be "200, 201, 204" respectively
    And user "Brian" should see the following elements
      | /Shares/new/|


  Scenario: keep user shared file name same after one of recipient has renamed the file inside the Shares folder
    Given user "Alice" has uploaded file with content "foo" to "/sharefile.txt"
    And user "Alice" has shared file "/sharefile.txt" with user "Brian"
    And user "Alice" has shared file "/sharefile.txt" with user "Carol"
    When user "Carol" moves file "/Shares/sharefile.txt" to "/Shares/renamedsharefile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Carol" file "/Shares/renamedsharefile.txt" should exist
    And as "Alice" file "/sharefile.txt" should exist
    And as "Brian" file "/Shares/sharefile.txt" should exist


  Scenario: receiver renames a received share with share, read, change permissions inside the Shares folder
    Given user "Alice" has created folder "folderToShare"
    And user "Alice" has uploaded file with content "thisIsAFileInsideTheSharedFolder" to "/folderToShare/fileInside"
    And user "Alice" has shared folder "folderToShare" with user "Brian" with permissions "share,read,change"
    When user "Brian" moves folder "/Shares/folderToShare" to "/Shares/myFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" folder "/Shares/myFolder" should exist
    But as "Alice" folder "/Shares/myFolder" should not exist
    When user "Brian" moves file "/Shares/myFolder/fileInside" to "/Shares/myFolder/renamedFile" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" file "/Shares/myFolder/renamedFile" should exist
    And as "Alice" file "/folderToShare/renamedFile" should exist
    But as "Alice" file "/folderToShare/fileInside" should not exist


  Scenario: receiver tries to rename a received share with share, read permissions inside the Shares folder
    Given user "Alice" has created folder "folderToShare"
    And user "Alice" has uploaded file with content "thisIsAFileInsideTheSharedFolder" to "/folderToShare/fileInside"
    And user "Alice" has shared folder "folderToShare" with user "Brian" with permissions "share,read"
    When user "Brian" moves folder "/Shares/folderToShare" to "/Shares/myFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" folder "/Shares/myFolder" should exist
    But as "Alice" folder "/Shares/myFolder" should not exist
    When user "Brian" moves file "/Shares/myFolder/fileInside" to "/Shares/myFolder/renamedFile" using the WebDAV API
    Then the HTTP status code should be "403"
    And as "Brian" file "/Shares/myFolder/renamedFile" should not exist
    But as "Brian" file "Shares/myFolder/fileInside" should exist


  Scenario: receiver renames a received folder share to a different name on the same folder
    Given user "Alice" has created folder "PARENT"
    And user "Alice" has shared folder "PARENT" with user "Brian"
    When user "Brian" moves folder "/Shares/PARENT" to "/Shares/myFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" folder "/Shares/myFolder" should exist
    But as "Alice" folder "myFolder" should not exist


  Scenario: receiver renames a received file share to different name on the same folder
    Given user "Alice" has uploaded file "filesForUpload/textfile.txt" to "fileToShare.txt"
    And user "Alice" has shared file "fileToShare.txt" with user "Brian"
    When user "Brian" moves file "/Shares/fileToShare.txt" to "/Shares/newFile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" file "/Shares/newFile.txt" should exist
    But as "Alice" file "newFile.txt" should not exist


  Scenario: receiver renames a received file share to different name on the same folder for group sharing
    Given group "grp1" has been created
    And user "Brian" has been added to group "grp1"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "fileToShare.txt"
    And user "Alice" has shared file "fileToShare.txt" with group "grp1"
    When user "Brian" moves file "/Shares/fileToShare.txt" to "/Shares/newFile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" file "/Shares/newFile.txt" should exist
    But as "Alice" file "newFile.txt" should not exist


  Scenario: receiver renames a received folder share to different name on the same folder for group sharing
    Given group "grp1" has been created
    And user "Alice" has created folder "PARENT"
    And user "Brian" has been added to group "grp1"
    And user "Alice" has shared folder "PARENT" with group "grp1"
    When user "Brian" moves folder "/Shares/PARENT" to "/Shares/myFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" folder "/Shares/myFolder" should exist
    But as "Alice" folder "myFolder" should not exist


  Scenario: receiver renames a received file share with read,update,share permissions inside the Shares folder in group sharing
    Given group "grp1" has been created
    And user "Brian" has been added to group "grp1"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "fileToShare.txt"
    And user "Alice" has shared file "fileToShare.txt" with group "grp1" with permissions "read,update,share"
    When user "Brian" moves folder "/Shares/fileToShare.txt" to "/Shares/newFile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" file "/Shares/newFile.txt" should exist
    But as "Alice" file "/Shares/newFile.txt" should not exist


  Scenario: receiver renames a received folder share with share, read, change permissions inside the Shares folder in group sharing
    Given group "grp1" has been created
    And user "Alice" has created folder "PARENT"
    And user "Brian" has been added to group "grp1"
    And user "Alice" has shared folder "PARENT" with group "grp1" with permissions "share,read,change"
    When user "Brian" moves folder "/Shares/PARENT" to "/Shares/myFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" folder "/Shares/myFolder" should exist
    But as "Alice" folder "/Shares/myFolder" should not exist


  Scenario: receiver renames a received file share with share, read permissions inside the Shares folder in group sharing)
    Given group "grp1" has been created
    And user "Brian" has been added to group "grp1"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "fileToShare.txt"
    And user "Alice" has shared file "fileToShare.txt" with group "grp1" with permissions "share,read"
    When user "Brian" moves file "/Shares/fileToShare.txt" to "/Shares/newFile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" file "/Shares/newFile.txt" should exist
    But as "Alice" file "/Shares/newFile.txt" should not exist


  Scenario: receiver renames a received folder share with share, read permissions inside the Shares folder in group sharing
    Given group "grp1" has been created
    And user "Alice" has created folder "PARENT"
    And user "Brian" has been added to group "grp1"
    And user "Alice" has shared folder "PARENT" with group "grp1" with permissions "share,read"
    When user "Brian" moves folder "/Shares/PARENT" to "/Shares/myFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Brian" folder "/Shares/myFolder" should exist
    But as "Alice" folder "/Shares/myFolder" should not exist

  @issue-2141
  Scenario Outline: receiver renames a received folder share to name with special characters in group sharing
    Given group "grp1" has been created
    And user "Carol" has been added to group "grp1"
    And user "Alice" has created folder "<sharer_folder>"
    And user "Alice" has created folder "<group_folder>"
    And user "Alice" has shared folder "<sharer_folder>" with user "Brian"
    When user "Brian" moves folder "/Shares/<sharer_folder>" to "/Shares/<receiver_folder>" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" folder "<receiver_folder>" should not exist
    And as "Brian" folder "/Shares/<receiver_folder>" should exist
    When user "Alice" shares folder "<group_folder>" with group "grp1" using the sharing API
    And user "Carol" moves folder "/Shares/<group_folder>" to "/Shares/<receiver_folder>" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" folder "<receiver_folder>" should not exist
    But as "Carol" folder "/Shares/<receiver_folder>" should exist
    Examples:
      | sharer_folder | group_folder    | receiver_folder |
      | ?abc=oc #     | ?abc=oc g%rp#   | # oc?test=oc&a  |
      | @a#8a=b?c=d   | @a#8a=b?c=d grp | ?a#8 a=b?c=d    |

  @issue-2141
  Scenario Outline: receiver renames a received file share to name with special characters with share, read, change permissions in group sharing
    Given group "grp1" has been created
    And user "Carol" has been added to group "grp1"
    And user "Alice" has created folder "<sharer_folder>"
    And user "Alice" has created folder "<group_folder>"
    And user "Alice" has uploaded file with content "thisIsAFileInsideTheSharedFolder" to "/<sharer_folder>/fileInside"
    And user "Alice" has uploaded file with content "thisIsAFileInsideTheSharedFolder" to "/<group_folder>/fileInside"
    And user "Alice" has shared folder "<sharer_folder>" with user "Brian" with permissions "share,read,change"
    When user "Brian" moves folder "/Shares/<sharer_folder>/fileInside" to "/Shares/<sharer_folder>/<receiver_file>" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" file "<sharer_folder>/<receiver_file>" should exist
    And as "Brian" file "/Shares/<sharer_folder>/<receiver_file>" should exist
    When user "Alice" shares folder "<group_folder>" with group "grp1" with permissions "share,read,change" using the sharing API
    And user "Carol" moves folder "/Shares/<group_folder>/fileInside" to "/Shares/<group_folder>/<receiver_file>" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" file "<group_folder>/<receiver_file>" should exist
    And as "Carol" file "/Shares/<group_folder>/<receiver_file>" should exist
    Examples:
      | sharer_folder | group_folder    | receiver_file  |
      | ?abc=oc #     | ?abc=oc g%rp#   | # oc?test=oc&a |
      | @a#8a=b?c=d   | @a#8a=b?c=d grp | ?a#8 a=b?c=d   |
