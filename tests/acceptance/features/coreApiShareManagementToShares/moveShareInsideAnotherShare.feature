@skipOnReva
Feature: moving a share inside another share
  As a user
  I want to move a shared resource inside another shared resource
  So that I have full flexibility when managing resources

  Background:
    Given using OCS API version "1"
    And these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
    And user "Alice" has created folder "folderA"
    And user "Alice" has created folder "folderB"
    And user "Alice" has uploaded file with content "text A" to "/folderA/fileA.txt"
    And user "Alice" has uploaded file with content "text B" to "/folderB/fileB.txt"
    And user "Alice" has shared folder "folderA" with user "Brian"
    And user "Alice" has shared folder "folderB" with user "Brian"


  Scenario: share receiver cannot move a whole share inside another share
    When user "Brian" moves folder "Shares/folderB" to "Shares/folderA/folderB" using the WebDAV API
    Then the HTTP status code should be "403"
    And as "Alice" folder "/folderB" should exist
    And as "Brian" folder "/Shares/folderB" should exist
    And as "Alice" file "/folderB/fileB.txt" should exist
    And as "Brian" file "/Shares/folderB/fileB.txt" should exist


  Scenario: share owner moves a whole share inside another share
    When user "Alice" moves folder "folderB" to "folderA/folderB" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" folder "/folderB" should not exist
    And as "Alice" folder "/folderA/folderB" should exist
    And as "Brian" folder "/Shares/folderB" should exist
    And as "Alice" file "/folderA/folderB/fileB.txt" should exist
    And as "Brian" file "/Shares/folderA/folderB/fileB.txt" should exist
    And as "Brian" file "/Shares/folderB/fileB.txt" should exist


  Scenario: share receiver moves a local folder inside a received share (local folder does not have a share in it)
    Given user "Brian" has created folder "localFolder"
    And user "Brian" has created folder "localFolder/subFolder"
    And user "Brian" has uploaded file with content "local text" to "/localFolder/localFile.txt"
    When user "Brian" moves folder "localFolder" to "Shares/folderA/localFolder" using the WebDAV API
    Then the HTTP status code should be "201"
    And as "Alice" folder "/folderA/localFolder" should exist
    And as "Brian" folder "/Shares/folderA/localFolder" should exist
    And as "Alice" folder "/folderA/localFolder/subFolder" should exist
    And as "Brian" folder "/Shares/folderA/localFolder/subFolder" should exist
    And as "Alice" file "/folderA/localFolder/localFile.txt" should exist
    And as "Brian" file "/Shares/folderA/localFolder/localFile.txt" should exist


  Scenario: share receiver tries to move a whole share inside a local folder
    Given user "Brian" has created folder "localFolder"
    And user "Brian" has uploaded file with content "local text" to "/localFolder/localFile.txt"
    # On oCIS you cannot move received shares out of the "Shares" folder
    When user "Brian" moves folder "Shares/folderB" to "localFolder/folderB" using the WebDAV API
    Then the HTTP status code should be "403"
    And as "Alice" file "/folderB/fileB.txt" should exist
    And as "Brian" file "/Shares/folderB/fileB.txt" should exist
