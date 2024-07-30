techupskills: @terminal How do I create a new branch for my GitHub repository?

GitHub Copilot: To create a new branch for your GitHub repository, you can use the following steps in the terminal:

1. **Open Terminal**: Open your terminal in Visual Studio Code or any terminal application.

2. **Navigate to Your Repository**: Ensure you are in the root directory of your GitHub repository.

3. **Create a New Branch**: Use the `git checkout` command with the `-b` option to create and switch to a new branch. Replace `new-branch-name` with your desired branch name.

```sh
git checkout -b new-branch-name
```

4. **Push the New Branch to GitHub**: To make the new branch available on GitHub, push it to the remote repository.

```sh
git push -u origin new-branch-name
```

### Example

```sh
git checkout -b feature-branch
git push -u origin feature-branch
```

This will create a new branch named `feature-branch` and push it to the remote repository, setting up tracking between your local branch and the remote branch.
