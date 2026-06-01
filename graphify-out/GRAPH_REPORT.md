# Graph Report - .  (2026-06-01)

## Corpus Check
- Corpus is ~23,906 words - fits in a single context window. You may not need a graph.

## Summary
- 283 nodes · 540 edges · 34 communities detected
- Extraction: 54% EXTRACTED · 46% INFERRED · 0% AMBIGUOUS · INFERRED: 247 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Test Suite|Test Suite]]
- [[_COMMUNITY_CLI Commands Core|CLI Commands Core]]
- [[_COMMUNITY_Remote Repository Types|Remote Repository Types]]
- [[_COMMUNITY_Command Handlers|Command Handlers]]
- [[_COMMUNITY_Getter Pipeline|Getter Pipeline]]
- [[_COMMUNITY_Migrate & Remove|Migrate & Remove]]
- [[_COMMUNITY_MigrateRemove Internals|Migrate/Remove Internals]]
- [[_COMMUNITY_GitHub Repository & Logger|GitHub Repository & Logger]]
- [[_COMMUNITY_Command Runner Util|Command Runner Util]]
- [[_COMMUNITY_Create Command|Create Command]]
- [[_COMMUNITY_Get Command|Get Command]]
- [[_COMMUNITY_VCS Operations|VCS Operations]]
- [[_COMMUNITY_Windows Symlink Tests|Windows Symlink Tests]]
- [[_COMMUNITY_Windows Path Helpers|Windows Path Helpers]]
- [[_COMMUNITY_Command Documentation|Command Documentation]]
- [[_COMMUNITY_CloneUpdate Args|Clone/Update Args]]
- [[_COMMUNITY_Git Worktree|Git Worktree]]
- [[_COMMUNITY_App Entry Point|App Entry Point]]
- [[_COMMUNITY_Shared Utilities|Shared Utilities]]
- [[_COMMUNITY_Input Scanner|Input Scanner]]
- [[_COMMUNITY_VCS Backend Interface|VCS Backend Interface]]
- [[_COMMUNITY_Subversion Backend|Subversion Backend]]
- [[_COMMUNITY_Mercurial Backend|Mercurial Backend]]
- [[_COMMUNITY_Bazaar Backend|Bazaar Backend]]
- [[_COMMUNITY_Go Import Detection|Go Import Detection]]
- [[_COMMUNITY_Linked Git Dir Check|Linked Git Dir Check]]
- [[_COMMUNITY_Main Repo Resolution|Main Repo Resolution]]
- [[_COMMUNITY_Get Command Def|Get Command Def]]
- [[_COMMUNITY_List Command Def|List Command Def]]
- [[_COMMUNITY_Root Command Def|Root Command Def]]
- [[_COMMUNITY_Create Command Def|Create Command Def]]
- [[_COMMUNITY_Local Repository Type|Local Repository Type]]
- [[_COMMUNITY_Getter Type|Getter Type]]
- [[_COMMUNITY_Run Error Type|Run Error Type]]

## God Nodes (most connected - your core abstractions)
1. `Run()` - 43 edges
2. `newTempDir()` - 25 edges
3. `newApp()` - 23 edges
4. `setEnv()` - 18 edges
5. `capture()` - 15 edges
6. `newURL()` - 15 edges
7. `Log()` - 13 edges
8. `doRm()` - 11 edges
9. `localRepositoryRoots()` - 11 edges
10. `TestRmWorktree()` - 10 edges

## Surprising Connections (you probably didn't know these)
- `newTempDir()` --calls--> `toFullPath()`  [INFERRED]
  helpers_test.go → helpers_unix.go
- `doGet command handler` --semantically_similar_to--> `doList command handler`  [INFERRED] [semantically similar]
  cmd_get.go → cmd_list.go
- `doList command handler` --semantically_similar_to--> `doRoot command handler`  [INFERRED] [semantically similar]
  cmd_list.go → cmd_root.go
- `doRm command handler` --semantically_similar_to--> `doMigrate command handler`  [INFERRED] [semantically similar]
  cmd_rm.go → cmd_migrate.go
- `TestNewRemoteRepository()` --calls--> `Run()`  [INFERRED]
  remote_repository_test.go → cmdutil/run.go

## Hyperedges (group relationships)
- **** — remote_repository_GitHubRepository, remote_repository_GitHubGistRepository, remote_repository_DarksHubRepository, remote_repository_NestPijulRepository, remote_repository_CodeCommitRepository, remote_repository_ChiselRepository, remote_repository_OtherRepository [EXTRACTED 1.00]
- **** — vcs_GitBackend, vcs_SubversionBackend, vcs_GitsvnBackend, vcs_MercurialBackend, vcs_DarcsBackend, vcs_PijulBackend, vcs_FossilBackend, vcs_BazaarBackend [EXTRACTED 1.00]
- **** — cmd_get_doGet, cmd_list_doList, cmd_root_doRoot, cmd_create_doCreate [EXTRACTED 1.00]
- **** — url_newURL, local_repository_LocalRepositoryFromURL, local_repository_localRepositoryRoots [INFERRED 0.85]
- **** — cmd_rm_doRm, cmd_migrate_doMigrate, cmd_rm_confirm [INFERRED 0.80]

## Communities

### Community 0 - "Test Suite"
Cohesion: 0.12
Nodes (41): Run(), TestDoCreate(), TestBareLook(), TestCommandGet(), TestCommandGet_gotMessage(), TestCommandGet_printPath(), TestDoGet_bulk(), equalPathLines() (+33 more)

### Community 1 - "CLI Commands Core"
Cohesion: 0.11
Nodes (23): doList(), doRoot(), evalSymlinks(), toFullPath(), findVCSBackend(), getHome(), getRoot(), LocalRepositoryFromFullPath() (+15 more)

### Community 2 - "Remote Repository Types"
Cohesion: 0.08
Nodes (11): ChiselRepository, CodeCommitRepository, DarksHubRepository, GitHubGistRepository, detectGoImport(), detectVCSAndRepoURL(), TestDetectVCSAndRepoURL(), metaImport (+3 more)

### Community 3 - "Command Handlers"
Cohesion: 0.1
Nodes (23): doCreate command handler, doGet command handler, getter struct, look function, lookByLocalRepository function, doList command handler, doRoot command handler, detectGoImport function (+15 more)

### Community 4 - "Getter Pipeline"
Cohesion: 0.18
Nodes (12): getInfo, getter, detectLocalRepoRoot(), getRepoLock(), TestDetectLocalRepoRoot(), convertGitURLHTTPToSSH(), detectUserName(), fillUsernameToPath() (+4 more)

### Community 5 - "Migrate & Remove"
Cohesion: 0.15
Nodes (17): doMigrate command handler, moveDir (cross-device safe), confirm prompt, doRm command handler, commandMigrate definition, commandRm definition, getter.get method, getter.getRemoteRepository (+9 more)

### Community 6 - "Migrate/Remove Internals"
Cohesion: 0.26
Nodes (11): doMigrate(), moveDir(), confirm(), doRm(), hasLinkedWorktrees(), isLinkedGitDir(), isNotADirectory(), isWorktreeGitDir() (+3 more)

### Community 7 - "GitHub Repository & Logger"
Cohesion: 0.24
Nodes (7): GitHubRepository, init(), Log(), Logf(), selectLogger(), SetOutput(), TestLog()

### Community 8 - "Command Runner Util"
Cohesion: 0.24
Nodes (10): RunCommand(), RunInDir(), RunInDirSilently(), RunSilently(), TestRun(), TestRunInDir(), TestRunInDirSilently(), TestRunSilently() (+2 more)

### Community 9 - "Create Command"
Cohesion: 0.31
Nodes (8): doCreate(), isNotExistOrEmpty(), mustParseURL(), TestNewLocalRepository(), NewRemoteRepository(), TestNewRemoteRepository(), TestNewRemoteRepository_error(), TestNewRemoteRepository_vcs_error()

### Community 10 - "Get Command"
Cohesion: 0.33
Nodes (7): detectShell(), doGet(), look(), lookByLocalRepository(), TestLook(), scanner, sliceScanner

### Community 11 - "VCS Operations"
Cohesion: 0.29
Nodes (4): replaceOnce(), svnBase(), VCSBackend, vcsGetOption

### Community 12 - "Windows Symlink Tests"
Cohesion: 0.67
Nodes (3): createLink(), Test_evalSymlinks(), testEvalSymlinksMode

### Community 13 - "Windows Path Helpers"
Cohesion: 0.67
Nodes (2): evalSymlinks(), filepathSplitAll()

### Community 14 - "Command Documentation"
Cohesion: 0.67
Nodes (3): commandDoc, init(), mkCommandsTemplate()

### Community 15 - "Clone/Update Args"
Cohesion: 0.67
Nodes (2): _cloneArgs, _updateArgs

### Community 16 - "Git Worktree"
Cohesion: 0.67
Nodes (3): hasLinkedWorktrees function, listLinkedWorktreePaths function, repairWorktreeBackPointers function

### Community 17 - "App Entry Point"
Cohesion: 0.67
Nodes (3): commands slice, main entrypoint, newApp (CLI app factory)

### Community 18 - "Shared Utilities"
Cohesion: 0.67
Nodes (3): cmdutil.Run, cmdutil.RunCommand, logger.Log function

### Community 19 - "Input Scanner"
Cohesion: 1.0
Nodes (1): scanner interface

### Community 20 - "VCS Backend Interface"
Cohesion: 1.0
Nodes (1): VCSBackend struct

### Community 21 - "Subversion Backend"
Cohesion: 1.0
Nodes (1): SubversionBackend VCS

### Community 22 - "Mercurial Backend"
Cohesion: 1.0
Nodes (1): MercurialBackend VCS

### Community 23 - "Bazaar Backend"
Cohesion: 1.0
Nodes (1): BazaarBackend VCS

### Community 24 - "Go Import Detection"
Cohesion: 1.0
Nodes (1): metaImport struct

### Community 25 - "Linked Git Dir Check"
Cohesion: 1.0
Nodes (1): isLinkedGitDir function

### Community 26 - "Main Repo Resolution"
Cohesion: 1.0
Nodes (1): resolveMainRepoDir function

### Community 27 - "Get Command Def"
Cohesion: 1.0
Nodes (1): commandGet definition

### Community 28 - "List Command Def"
Cohesion: 1.0
Nodes (1): commandList definition

### Community 29 - "Root Command Def"
Cohesion: 1.0
Nodes (1): commandRoot definition

### Community 30 - "Create Command Def"
Cohesion: 1.0
Nodes (1): commandCreate definition

### Community 31 - "Local Repository Type"
Cohesion: 1.0
Nodes (1): LocalRepository struct

### Community 32 - "Getter Type"
Cohesion: 1.0
Nodes (1): getter struct

### Community 33 - "Run Error Type"
Cohesion: 1.0
Nodes (1): RunError type

## Knowledge Gaps
- **47 isolated node(s):** `scanner`, `testEvalSymlinksMode`, `_cloneArgs`, `_updateArgs`, `metaImport` (+42 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Windows Path Helpers`** (4 nodes): `evalSymlinks()`, `filepathSplitAll()`, `toFullPath()`, `helpers_windows.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Clone/Update Args`** (3 nodes): `commands_test.go`, `_cloneArgs`, `_updateArgs`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Input Scanner`** (1 nodes): `scanner interface`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `VCS Backend Interface`** (1 nodes): `VCSBackend struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Subversion Backend`** (1 nodes): `SubversionBackend VCS`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Mercurial Backend`** (1 nodes): `MercurialBackend VCS`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Bazaar Backend`** (1 nodes): `BazaarBackend VCS`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Go Import Detection`** (1 nodes): `metaImport struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Linked Git Dir Check`** (1 nodes): `isLinkedGitDir function`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Main Repo Resolution`** (1 nodes): `resolveMainRepoDir function`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Get Command Def`** (1 nodes): `commandGet definition`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `List Command Def`** (1 nodes): `commandList definition`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Root Command Def`** (1 nodes): `commandRoot definition`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Create Command Def`** (1 nodes): `commandCreate definition`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Local Repository Type`** (1 nodes): `LocalRepository struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Getter Type`** (1 nodes): `getter struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Run Error Type`** (1 nodes): `RunError type`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Run()` connect `Test Suite` to `CLI Commands Core`, `Getter Pipeline`, `GitHub Repository & Logger`, `Command Runner Util`, `Create Command`, `Get Command`, `Windows Symlink Tests`?**
  _High betweenness centrality (0.201) - this node is a cross-community bridge._
- **Why does `newURL()` connect `Getter Pipeline` to `Test Suite`, `CLI Commands Core`, `Migrate/Remove Internals`, `GitHub Repository & Logger`, `Create Command`, `Get Command`?**
  _High betweenness centrality (0.082) - this node is a cross-community bridge._
- **Why does `Log()` connect `GitHub Repository & Logger` to `Test Suite`, `CLI Commands Core`, `Remote Repository Types`, `Getter Pipeline`, `Migrate/Remove Internals`, `Command Runner Util`?**
  _High betweenness centrality (0.069) - this node is a cross-community bridge._
- **Are the 41 inferred relationships involving `Run()` (e.g. with `TestNewRemoteRepository()` and `TestNewRemoteRepository_vcs_error()`) actually correct?**
  _`Run()` has 41 INFERRED edges - model-reasoned connections that need verification._
- **Are the 24 inferred relationships involving `newTempDir()` (e.g. with `withFakeGitBackend()` and `TestRmCommand()`) actually correct?**
  _`newTempDir()` has 24 INFERRED edges - model-reasoned connections that need verification._
- **Are the 21 inferred relationships involving `newApp()` (e.g. with `TestRmCommand()` and `TestRmDryRunCommand()`) actually correct?**
  _`newApp()` has 21 INFERRED edges - model-reasoned connections that need verification._
- **Are the 17 inferred relationships involving `setEnv()` (e.g. with `TestRmCommand()` and `TestRmDryRunCommand()`) actually correct?**
  _`setEnv()` has 17 INFERRED edges - model-reasoned connections that need verification._