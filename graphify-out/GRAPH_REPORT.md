# Graph Report - .  (2026-06-01)

## Corpus Check
- 42 files · ~27,935 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 308 nodes · 599 edges · 31 communities detected
- Extraction: 54% EXTRACTED · 46% INFERRED · 0% AMBIGUOUS · INFERRED: 275 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- [[_COMMUNITY_Command Tests|Command Tests]]
- [[_COMMUNITY_CLI Command Impl|CLI Command Impl]]
- [[_COMMUNITY_Remote Repository Types|Remote Repository Types]]
- [[_COMMUNITY_Migrate & Remove|Migrate & Remove]]
- [[_COMMUNITY_GetClone Pipeline|Get/Clone Pipeline]]
- [[_COMMUNITY_Get Command Logic|Get Command Logic]]
- [[_COMMUNITY_Getter & Branch|Getter & Branch]]
- [[_COMMUNITY_Migrate & Remove Ops|Migrate & Remove Ops]]
- [[_COMMUNITY_Cmdutil Run Layer|Cmdutil Run Layer]]
- [[_COMMUNITY_VCS Backend Helpers|VCS Backend Helpers]]
- [[_COMMUNITY_Subprocess Execution|Subprocess Execution]]
- [[_COMMUNITY_Windows Helpers|Windows Helpers]]
- [[_COMMUNITY_Command Registry|Command Registry]]
- [[_COMMUNITY_CloneUpdate Args|Clone/Update Args]]
- [[_COMMUNITY_Main Test|Main Test]]
- [[_COMMUNITY_URL Parsing|URL Parsing]]
- [[_COMMUNITY_Local Repository Walker|Local Repository Walker]]
- [[_COMMUNITY_Worktree Management|Worktree Management]]
- [[_COMMUNITY_Go Import Resolution|Go Import Resolution]]
- [[_COMMUNITY_Logger|Logger]]
- [[_COMMUNITY_Remote Repo Tests|Remote Repo Tests]]
- [[_COMMUNITY_Helpers Unix|Helpers Unix]]
- [[_COMMUNITY_Worktree Clone Tests|Worktree Clone Tests]]
- [[_COMMUNITY_URL Tests|URL Tests]]
- [[_COMMUNITY_List Command|List Command]]
- [[_COMMUNITY_Root Command|Root Command]]
- [[_COMMUNITY_Documentation|Documentation]]
- [[_COMMUNITY_Create Command|Create Command]]
- [[_COMMUNITY_VCS Backends|VCS Backends]]
- [[_COMMUNITY_Git Operations|Git Operations]]
- [[_COMMUNITY_Repository Interface|Repository Interface]]

## God Nodes (most connected - your core abstractions)
1. `Run()` - 48 edges
2. `newTempDir()` - 29 edges
3. `newApp()` - 23 edges
4. `setEnv()` - 18 edges
5. `newURL()` - 16 edges
6. `capture()` - 15 edges
7. `Log()` - 14 edges
8. `doRm()` - 11 edges
9. `LocalRepositoryFromURL()` - 11 edges
10. `localRepositoryRoots()` - 11 edges

## Surprising Connections (you probably didn't know these)
- `detectDefaultBranch` --semantically_similar_to--> `doMigrateWorktree Command`  [INFERRED] [semantically similar]
  branch.go → cmd_migrate_worktree.go
- `newTempDir()` --calls--> `toFullPath()`  [INFERRED]
  helpers_test.go → helpers_unix.go
- `TestMain()` --calls--> `Run()`  [INFERRED]
  main_test.go → cmdutil/run.go
- `doGet Command` --semantically_similar_to--> `doList command handler`  [INFERRED] [semantically similar]
  cmd_get.go → cmd_list.go
- `doList command handler` --semantically_similar_to--> `doRoot command handler`  [INFERRED] [semantically similar]
  cmd_list.go → cmd_root.go

## Hyperedges (group relationships)
- **** — cmd_get_doGet, cmd_list_doList, cmd_root_doRoot, cmd_create_doCreate [EXTRACTED 1.00]
- **Worktree Management Flow** — worktree_isLinkedGitDir, worktree_isWorktreeGitDir, worktree_hasLinkedWorktrees, worktree_listLinkedWorktreePaths, worktree_resolveMainRepoDir, worktree_repairWorktreeBackPointers [EXTRACTED 0.90]
- **Remote Repository Type Dispatch** — remote_repository_NewRemoteRepository, remote_repository_GitHubRepository, remote_repository_OtherRepository, remote_repository_CodeCommitRepository, remote_repository_ChiselRepository [EXTRACTED 1.00]
- **VCS Clone Worktree Layout** — vcs_GitBackend, branch_detectDefaultBranch, cmd_migrate_worktree_doMigrateWorktree [INFERRED 0.80]
- **CLI Command Registration and Dispatch** — main_newApp, commands_commands, commands_commandGet, commands_commandList, commands_commandMigrate, commands_commandMigrateWorktree [EXTRACTED 1.00]
- **Repository Resolution Flow** — getter_getRemoteRepository, local_repository_LocalRepositoryFromURL, local_repository_getRoot, local_repository_walkLocalRepositories [EXTRACTED 0.90]
- **Subprocess Execution Layer** — cmdutil_Run, cmdutil_RunInDir, cmdutil_RunCommand, cmdutil_CommandRunner [EXTRACTED 1.00]

## Communities

### Community 0 - "Command Tests"
Cohesion: 0.12
Nodes (46): Run(), TestDoCreate(), TestBareLook(), TestCommandGet(), TestCommandGet_gotMessage(), TestCommandGet_printPath(), TestDoGet_bulk(), TestLook() (+38 more)

### Community 1 - "CLI Command Impl"
Cohesion: 0.07
Nodes (35): doCreate(), isNotExistOrEmpty(), doList(), doMigrateWorktree(), doRoot(), samePaths(), evalSymlinks(), toFullPath() (+27 more)

### Community 2 - "Remote Repository Types"
Cohesion: 0.08
Nodes (11): ChiselRepository, CodeCommitRepository, DarksHubRepository, GitHubGistRepository, detectGoImport(), detectVCSAndRepoURL(), TestDetectVCSAndRepoURL(), metaImport (+3 more)

### Community 3 - "Migrate & Remove"
Cohesion: 0.09
Nodes (31): doMigrate, moveDir, doMigrateWorktree Command, confirm prompt, doRm Command, commandGet (get/clone command), commandList (list command), commandMigrate (migrate command) (+23 more)

### Community 4 - "Get/Clone Pipeline"
Cohesion: 0.11
Nodes (24): detectDefaultBranch, doCreate command handler, doGet Command, getter struct, look function, lookByLocalRepository function, doList command handler, doRoot command handler (+16 more)

### Community 5 - "Get Command Logic"
Cohesion: 0.15
Nodes (13): detectShell(), doGet(), look(), lookByLocalRepository(), GitHubRepository, scanner, sliceScanner, init() (+5 more)

### Community 6 - "Getter & Branch"
Cohesion: 0.14
Nodes (14): detectDefaultBranch(), getInfo, getter, detectLocalRepoRoot(), getRepoLock(), TestDetectLocalRepoRoot(), convertGitURLHTTPToSSH(), detectUserName() (+6 more)

### Community 7 - "Migrate & Remove Ops"
Cohesion: 0.25
Nodes (13): doMigrate(), moveDir(), confirm(), doRm(), hasLinkedWorktrees(), hasWorktreeEntries(), isLinkedGitDir(), isNotADirectory() (+5 more)

### Community 8 - "Cmdutil Run Layer"
Cohesion: 0.24
Nodes (10): RunCommand(), RunInDir(), RunInDirSilently(), RunSilently(), TestRun(), TestRunInDir(), TestRunInDirSilently(), TestRunSilently() (+2 more)

### Community 9 - "VCS Backend Helpers"
Cohesion: 0.29
Nodes (4): replaceOnce(), svnBase(), VCSBackend, vcsGetOption

### Community 10 - "Subprocess Execution"
Cohesion: 0.4
Nodes (5): cmdutil.CommandRunner, cmdutil.Run, cmdutil.RunCommand, cmdutil.RunInDir, logger.Log function

### Community 11 - "Windows Helpers"
Cohesion: 0.67
Nodes (2): evalSymlinks(), filepathSplitAll()

### Community 12 - "Command Registry"
Cohesion: 0.67
Nodes (3): commandDoc, init(), mkCommandsTemplate()

### Community 13 - "Clone/Update Args"
Cohesion: 0.67
Nodes (2): _cloneArgs, _updateArgs

### Community 14 - "Main Test"
Cohesion: 1.0
Nodes (1): TestMain()

### Community 15 - "URL Parsing"
Cohesion: 1.0
Nodes (1): scanner interface

### Community 16 - "Local Repository Walker"
Cohesion: 1.0
Nodes (1): VCSBackend Struct

### Community 17 - "Worktree Management"
Cohesion: 1.0
Nodes (1): SubversionBackend VCS

### Community 18 - "Go Import Resolution"
Cohesion: 1.0
Nodes (1): MercurialBackend VCS

### Community 19 - "Logger"
Cohesion: 1.0
Nodes (1): BazaarBackend VCS

### Community 20 - "Remote Repo Tests"
Cohesion: 1.0
Nodes (1): metaImport struct

### Community 21 - "Helpers Unix"
Cohesion: 1.0
Nodes (1): resolveMainRepoDir

### Community 22 - "Worktree Clone Tests"
Cohesion: 1.0
Nodes (1): commandRoot definition

### Community 23 - "URL Tests"
Cohesion: 1.0
Nodes (1): commandCreate definition

### Community 24 - "List Command"
Cohesion: 1.0
Nodes (1): LocalRepository struct

### Community 25 - "Root Command"
Cohesion: 1.0
Nodes (1): getter struct

### Community 26 - "Documentation"
Cohesion: 1.0
Nodes (1): RunError type

### Community 27 - "Create Command"
Cohesion: 1.0
Nodes (1): getGitRemoteURL

### Community 28 - "VCS Backends"
Cohesion: 1.0
Nodes (1): vcsGetOption

### Community 29 - "Git Operations"
Cohesion: 1.0
Nodes (1): ghq - manage remote repository clones

### Community 30 - "Repository Interface"
Cohesion: 1.0
Nodes (1): AGENTS.md architecture notes

## Knowledge Gaps
- **51 isolated node(s):** `scanner`, `testEvalSymlinksMode`, `_cloneArgs`, `_updateArgs`, `metaImport` (+46 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **Thin community `Windows Helpers`** (4 nodes): `evalSymlinks()`, `filepathSplitAll()`, `toFullPath()`, `helpers_windows.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Clone/Update Args`** (3 nodes): `commands_test.go`, `_cloneArgs`, `_updateArgs`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Main Test`** (2 nodes): `TestMain()`, `main_test.go`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `URL Parsing`** (1 nodes): `scanner interface`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Local Repository Walker`** (1 nodes): `VCSBackend Struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Worktree Management`** (1 nodes): `SubversionBackend VCS`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Go Import Resolution`** (1 nodes): `MercurialBackend VCS`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Logger`** (1 nodes): `BazaarBackend VCS`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Remote Repo Tests`** (1 nodes): `metaImport struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Helpers Unix`** (1 nodes): `resolveMainRepoDir`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Worktree Clone Tests`** (1 nodes): `commandRoot definition`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `URL Tests`** (1 nodes): `commandCreate definition`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `List Command`** (1 nodes): `LocalRepository struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Root Command`** (1 nodes): `getter struct`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Documentation`** (1 nodes): `RunError type`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Create Command`** (1 nodes): `getGitRemoteURL`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `VCS Backends`** (1 nodes): `vcsGetOption`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Git Operations`** (1 nodes): `ghq - manage remote repository clones`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.
- **Thin community `Repository Interface`** (1 nodes): `AGENTS.md architecture notes`
  Too small to be a meaningful cluster - may be noise or needs more connections extracted.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Run()` connect `Command Tests` to `CLI Command Impl`, `Get Command Logic`, `Getter & Branch`, `Cmdutil Run Layer`, `Main Test`?**
  _High betweenness centrality (0.196) - this node is a cross-community bridge._
- **Why does `newURL()` connect `Getter & Branch` to `Command Tests`, `CLI Command Impl`, `Get Command Logic`, `Migrate & Remove Ops`?**
  _High betweenness centrality (0.076) - this node is a cross-community bridge._
- **Why does `Log()` connect `Get Command Logic` to `Command Tests`, `CLI Command Impl`, `Remote Repository Types`, `Getter & Branch`, `Migrate & Remove Ops`, `Cmdutil Run Layer`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Are the 46 inferred relationships involving `Run()` (e.g. with `TestNewRemoteRepository()` and `TestNewRemoteRepository_vcs_error()`) actually correct?**
  _`Run()` has 46 INFERRED edges - model-reasoned connections that need verification._
- **Are the 28 inferred relationships involving `newTempDir()` (e.g. with `withFakeGitBackend()` and `TestRmCommand()`) actually correct?**
  _`newTempDir()` has 28 INFERRED edges - model-reasoned connections that need verification._
- **Are the 21 inferred relationships involving `newApp()` (e.g. with `TestRmCommand()` and `TestRmDryRunCommand()`) actually correct?**
  _`newApp()` has 21 INFERRED edges - model-reasoned connections that need verification._
- **Are the 17 inferred relationships involving `setEnv()` (e.g. with `TestRmCommand()` and `TestRmDryRunCommand()`) actually correct?**
  _`setEnv()` has 17 INFERRED edges - model-reasoned connections that need verification._