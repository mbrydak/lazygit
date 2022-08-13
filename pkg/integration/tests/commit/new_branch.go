package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/integration/components"
)

var NewBranch = components.NewIntegrationTest(components.NewIntegrationTestArgs{
	Description:  "Creating a new branch from a commit",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *components.Shell) {
		shell.
			EmptyCommit("commit 1").
			EmptyCommit("commit 2").
			EmptyCommit("commit 3")
	},
	Run: func(shell *components.Shell, input *components.Input, assert *components.Assert, keys config.KeybindingConfig) {
		assert.CommitCount(3)

		input.SwitchToCommitsWindow()
		assert.CurrentViewName("commits")
		input.NextItem()

		input.PressKeys(keys.Universal.New)

		assert.CurrentViewName("confirmation")

		branchName := "my-branch-name"
		input.Type(branchName)
		input.Confirm()

		assert.CommitCount(2)
		assert.HeadCommitMessage("commit 2")
		assert.CurrentBranchName(branchName)
	},
})
