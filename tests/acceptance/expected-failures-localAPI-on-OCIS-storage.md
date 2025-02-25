## Scenarios from OCIS API tests that are expected to fail with OCIS storage

The expected failures in this file are from features in the owncloud/ocis repo.

#### [Downloading the archive of the resource (files | folder) using resource path is not possible](https://github.com/owncloud/ocis/issues/4637)

- [apiArchiver/downloadByPath.feature:25](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L25)
- [apiArchiver/downloadByPath.feature:26](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L26)
- [apiArchiver/downloadByPath.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L43)
- [apiArchiver/downloadByPath.feature:44](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L44)
- [apiArchiver/downloadByPath.feature:47](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L47)
- [apiArchiver/downloadByPath.feature:73](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L73)
- [apiArchiver/downloadByPath.feature:123](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L123)
- [apiArchiver/downloadByPath.feature:124](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadByPath.feature#L124)

### [Downloaded /Shares tar contains resource (files|folder) with leading / in Response](https://github.com/owncloud/ocis/issues/4636)

- [apiArchiver/downloadById.feature:125](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadById.feature#L125)
- [apiArchiver/downloadById.feature:126](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiArchiver/downloadById.feature#L126)

### [PROPFIND on accepted shares with identical names containing brackets exit with 404](https://github.com/owncloud/ocis/issues/4421)

- [apiSpacesShares/changingFilesShare.feature:14](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/changingFilesShare.feature#L14)

### [Shared mount folder gets deleted when overwritten by a file from personal space](https://github.com/owncloud/ocis/issues/7208)

- [apiSpacesShares/copySpaces.feature:510](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/copySpaces.feature#L510)
- [apiSpacesShares/copySpaces.feature:523](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/copySpaces.feature#L523)

#### [PATCH request for TUS upload with wrong checksum gives incorrect response](https://github.com/owncloud/ocis/issues/1755)

- [apiSpacesShares/shareUploadTUS.feature:187](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareUploadTUS.feature#L187)
- [apiSpacesShares/shareUploadTUS.feature:201](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareUploadTUS.feature#L201)
- [apiSpacesShares/shareUploadTUS.feature:264](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSpacesShares/shareUploadTUS.feature#L264)

### [Settings service user can list other peoples assignments](https://github.com/owncloud/ocis/issues/5032)

- [apiAccountsHashDifficulty/assignRole.feature:27](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiAccountsHashDifficulty/assignRole.feature#L27)
- [apiAccountsHashDifficulty/assignRole.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiAccountsHashDifficulty/assignRole.feature#L28)
- [apiGraph/assignRole.feature:30](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/assignRole.feature#L30)
- [apiGraph/assignRole.feature:31](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/assignRole.feature#L31)
- [apiGraph/assignRole.feature:32](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/assignRole.feature#L32)

#### [Share lists deleted user as 'user'](https://github.com/owncloud/ocis/issues/903)

- [apiGraph/deleteGroup.feature:67](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L67)

#### [CORS headers are not identical with oC10 headers](https://github.com/owncloud/ocis/issues/5195)

- [apiCors/cors.feature:28](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L28)
- [apiCors/cors.feature:29](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L29)
- [apiCors/cors.feature:30](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L30)
- [apiCors/cors.feature:31](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L31)

#### [Requests with invalid credentials do not return CORS headers](https://github.com/owncloud/ocis/issues/5194)

- [apiCors/cors.feature:70](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L70)
- [apiCors/cors.feature:71](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiCors/cors.feature#L71)

#### [A User can get information of another user with Graph API](https://github.com/owncloud/ocis/issues/5125)

- [apiGraph/getUser.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L87)
- [apiGraph/getUser.feature:88](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L88)
- [apiGraph/getUser.feature:89](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L89)
- [apiGraph/getUser.feature:90](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L90)
- [apiGraph/getUser.feature:91](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L91)
- [apiGraph/getUser.feature:92](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L92)
- [apiGraph/getUser.feature:93](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L93)
- [apiGraph/getUser.feature:94](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L94)
- [apiGraph/getUser.feature:95](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L95)
- [apiGraph/getUser.feature:96](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L96)
- [apiGraph/getUser.feature:97](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L97)
- [apiGraph/getUser.feature:98](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L98)
- [apiGraph/getUser.feature:642](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L642)
- [apiGraph/getUser.feature:643](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L643)
- [apiGraph/getUser.feature:644](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L644)
- [apiGraph/getUser.feature:645](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L645)
- [apiGraph/getUser.feature:646](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L646)
- [apiGraph/getUser.feature:647](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L647)
- [apiGraph/getUser.feature:648](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L648)
- [apiGraph/getUser.feature:649](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L649)
- [apiGraph/getUser.feature:650](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L650)
- [apiGraph/getUser.feature:651](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L651)
- [apiGraph/getUser.feature:652](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L652)
- [apiGraph/getUser.feature:653](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getUser.feature#L653)

#### [Normal user can get expanded members information of a group](https://github.com/owncloud/ocis/issues/5604)

- [apiGraph/getGroup.feature:381](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L381)
- [apiGraph/getGroup.feature:382](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L382)
- [apiGraph/getGroup.feature:383](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L383)

#### [Same users can be added in a group multiple time](https://github.com/owncloud/ocis/issues/5702)

- [apiGraph/addUserToGroup.feature:285](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L285)

#### [API requests from an unauthorized user should return 403](https://github.com/owncloud/ocis/issues/5938)

- [apiGraph/addUserToGroup.feature:150](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L150)
- [apiGraph/addUserToGroup.feature:151](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L151)
- [apiGraph/addUserToGroup.feature:152](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L152)
- [apiGraph/addUserToGroup.feature:184](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L184)
- [apiGraph/addUserToGroup.feature:185](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L185)
- [apiGraph/addUserToGroup.feature:186](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L186)
- [apiGraph/createGroup.feature:42](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L42)
- [apiGraph/createGroup.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L43)
- [apiGraph/createGroup.feature:44](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/createGroup.feature#L44)
- [apiGraph/deleteGroup.feature:63](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L63)
- [apiGraph/deleteGroup.feature:62](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L62)
- [apiGraph/deleteGroup.feature:64](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/deleteGroup.feature#L64)
- [apiGraph/editGroup.feature:35](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editGroup.feature#L35)
- [apiGraph/editGroup.feature:34](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editGroup.feature#L34)
- [apiGraph/editGroup.feature:36](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/editGroup.feature#L36)
- [apiGraph/getGroup.feature:103](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L103)
- [apiGraph/getGroup.feature:104](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L104)
- [apiGraph/getGroup.feature:105](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/getGroup.feature#L105)
- [apiGraph/removeUserFromGroup.feature:191](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/removeUserFromGroup.feature#L191)
- [apiGraph/removeUserFromGroup.feature:192](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/removeUserFromGroup.feature#L192)
- [apiGraph/removeUserFromGroup.feature:193](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/removeUserFromGroup.feature#L193)

#### [API requests for a non-existent resources should return 404](https://github.com/owncloud/ocis/issues/5939)

- [apiGraph/addUserToGroup.feature:201](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L201)
- [apiGraph/addUserToGroup.feature:202](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L202)
- [apiGraph/addUserToGroup.feature:203](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L203)

### [Users are added in a group with wrong host in host-part of user](https://github.com/owncloud/ocis/issues/5871)

- [apiGraph/addUserToGroup.feature:369](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L369)
- [apiGraph/addUserToGroup.feature:383](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L383)

### [Adding the same user as multiple members in a single request results in listing the same user twice in the group](https://github.com/owncloud/ocis/issues/5855)

- [apiGraph/addUserToGroup.feature:420](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiGraph/addUserToGroup.feature#L420)

### [Shared file locking is not possible using different path](https://github.com/owncloud/ocis/issues/7599)

- [apiLocks/lockFiles.feature:179](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L179)
- [apiLocks/lockFiles.feature:180](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L180)
- [apiLocks/lockFiles.feature:181](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L181)
- [apiLocks/lockFiles.feature:280](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L280)
- [apiLocks/lockFiles.feature:281](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L281)
- [apiLocks/lockFiles.feature:282](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L282)
- [apiLocks/lockFiles.feature:323](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L323)
- [apiLocks/lockFiles.feature:324](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L324)
- [apiLocks/lockFiles.feature:325](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L325)
- [apiLocks/lockFiles.feature:326](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L326)
- [apiLocks/lockFiles.feature:327](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L327)
- [apiLocks/lockFiles.feature:328](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L328)
- [apiLocks/lockFiles.feature:346](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L346)
- [apiLocks/lockFiles.feature:347](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L347)
- [apiLocks/lockFiles.feature:348](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L348)
- [apiLocks/lockFiles.feature:349](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L349)
- [apiLocks/lockFiles.feature:350](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L350)
- [apiLocks/lockFiles.feature:351](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L351)
- [apiLocks/unlockFiles.feature:60](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L60)
- [apiLocks/unlockFiles.feature:61](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L61)
- [apiLocks/unlockFiles.feature:62](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L62)
- [apiLocks/unlockFiles.feature:151](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L151)
- [apiLocks/unlockFiles.feature:152](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L152)
- [apiLocks/unlockFiles.feature:153](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L153)
- [apiLocks/unlockFiles.feature:154](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L154)
- [apiLocks/unlockFiles.feature:155](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L155)
- [apiLocks/unlockFiles.feature:156](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L156)
- [apiLocks/unlockFiles.feature:173](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L173)
- [apiLocks/unlockFiles.feature:174](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L174)
- [apiLocks/unlockFiles.feature:175](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L175)
- [apiLocks/unlockFiles.feature:176](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L176)
- [apiLocks/unlockFiles.feature:177](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L177)
- [apiLocks/unlockFiles.feature:178](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L178)
- [apiLocks/unlockFiles.feature:195](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L195)
- [apiLocks/unlockFiles.feature:196](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L196)
- [apiLocks/unlockFiles.feature:197](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L197)
- [apiLocks/unlockFiles.feature:198](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L198)
- [apiLocks/unlockFiles.feature:199](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L199)
- [apiLocks/unlockFiles.feature:200](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L200)

#### [Trying to upload to a locked file gives 500](https://github.com/owncloud/ocis/issues/7638)

- [apiLocks/lockFiles.feature:299](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L299)
- [apiLocks/lockFiles.feature:300](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L300)
- [apiLocks/lockFiles.feature:301](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L301)
- [apiLocks/lockFiles.feature:302](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L302)
- [apiLocks/lockFiles.feature:303](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L303)
- [apiLocks/lockFiles.feature:304](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L304)
- [apiLocks/unlockFiles.feature:85](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L85)
- [apiLocks/unlockFiles.feature:86](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L86)
- [apiLocks/unlockFiles.feature:87](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L87)
- [apiLocks/unlockFiles.feature:88](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L88)
- [apiLocks/unlockFiles.feature:89](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L89)
- [apiLocks/unlockFiles.feature:90](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L90)
- [apiLocks/lockFiles.feature:388](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L388)
- [apiLocks/lockFiles.feature:389](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L389)
- [apiLocks/lockFiles.feature:390](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L390)
- [apiLocks/lockFiles.feature:391](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L391)
- [apiLocks/lockFiles.feature:392](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L392)
- [apiLocks/lockFiles.feature:393](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L393)
- [apiLocks/lockFiles.feature:429](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L429)
- [apiLocks/lockFiles.feature:430](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L430)
- [apiLocks/lockFiles.feature:431](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L431)
- [apiLocks/lockFiles.feature:432](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L432)
- [apiLocks/lockFiles.feature:433](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L433)
- [apiLocks/lockFiles.feature:434](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L434)

#### [Folders can be locked and locking works partially](https://github.com/owncloud/ocis/issues/7641)

- [apiLocks/lockFiles.feature:364](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L364)
- [apiLocks/lockFiles.feature:365](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L365)
- [apiLocks/lockFiles.feature:366](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L366)
- [apiLocks/lockFiles.feature:367](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L367)
- [apiLocks/lockFiles.feature:368](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L368)
- [apiLocks/lockFiles.feature:369](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L369)

### [Anonymous users can unlock a file shared to them through a public link if they get the lock token](https://github.com/owncloud/ocis/issues/7761)
- [apiLocks/unlockFiles.feature:40](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L40)
- [apiLocks/unlockFiles.feature:41](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L41)
- [apiLocks/unlockFiles.feature:42](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L42)
- [apiLocks/unlockFiles.feature:43](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L43)
- [apiLocks/unlockFiles.feature:44](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L44)
- [apiLocks/unlockFiles.feature:45](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L45)

### [Trying to unlock a shared file with sharer's lock token gives 500](https://github.com/owncloud/ocis/issues/7767)
- [apiLocks/unlockFiles.feature:107](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L107)
- [apiLocks/unlockFiles.feature:108](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L108)
- [apiLocks/unlockFiles.feature:109](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L109)
- [apiLocks/unlockFiles.feature:110](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L110)
- [apiLocks/unlockFiles.feature:111](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L111)
- [apiLocks/unlockFiles.feature:112](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L112)
- [apiLocks/unlockFiles.feature:129](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L129)
- [apiLocks/unlockFiles.feature:130](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L130)
- [apiLocks/unlockFiles.feature:131](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L131)
- [apiLocks/unlockFiles.feature:132](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L132)
- [apiLocks/unlockFiles.feature:133](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L133)
- [apiLocks/unlockFiles.feature:134](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/unlockFiles.feature#L134)

### [Anonymous user trying lock a file shared to them through a public link gives 405](https://github.com/owncloud/ocis/issues/7790)
- [apiLocks/lockFiles.feature:474](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L474)
- [apiLocks/lockFiles.feature:475](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L475)
- [apiLocks/lockFiles.feature:476](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L476)
- [apiLocks/lockFiles.feature:477](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L477)
- [apiLocks/lockFiles.feature:478](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L478)
- [apiLocks/lockFiles.feature:479](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L479)
- [apiLocks/lockFiles.feature:496](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L496)
- [apiLocks/lockFiles.feature:497](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L497)
- [apiLocks/lockFiles.feature:498](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L498)
- [apiLocks/lockFiles.feature:499](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L499)
- [apiLocks/lockFiles.feature:500](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L500)
- [apiLocks/lockFiles.feature:501](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L501)

### [anonymous user with viewer role in public link of a folder can lock a file inside it](https://github.com/owncloud/ocis/issues/7785)
- [apiLocks/lockFiles.feature:452](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L452)
- [apiLocks/lockFiles.feature:453](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L453)
- [apiLocks/lockFiles.feature:454](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L454)
- [apiLocks/lockFiles.feature:455](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L455)
- [apiLocks/lockFiles.feature:456](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L456)
- [apiLocks/lockFiles.feature:457](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiLocks/lockFiles.feature#L457)

### [blocksDownload link type is not implemented yet (sharing-ng)](https://github.com/owncloud/ocis/issues/7879)
- [apiSharingNg/createLink.feature:78](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSharingNg/createLink.feature#L78)
- [apiSharingNg/createLink.feature:147](https://github.com/owncloud/ocis/blob/master/tests/acceptance/features/apiSharingNg/createLink.feature#L147)

- Note: always have an empty line at the end of this file.
The bash script that processes this file requires that the last line has a newline on the end.
