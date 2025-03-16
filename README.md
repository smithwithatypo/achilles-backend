# Project Name

This is the backend for Achilles

## Getting Started

### Prerequisites

- Go (https://golang.org/dl/)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/smithwithatypo/achilles-backend.git
    ```
2. Navigate to the project directory:
    ```sh
    cd achilles-backend
    ```

### Running the Project

To run the project, use the following command:
```sh
go run server.go
```

## Git Branch Management

### Pruning Remote and Local Branches

To clean up your Git repository by removing references to deleted remote branches and then removing local branches:

1. Prune remote-tracking branches that no longer exist on the remote:
```bash
git fetch --prune
```

2. List local branches that have been merged into the current branch (to identify candidates for deletion):
```bash
git branch --merged
```

3. Delete a specific local branch:
```bash
git branch -d <branch-name>
```

4. Force delete a local branch (if it contains unmerged changes):
```bash
git branch -D <branch-name>
```

5. To list local branches that track deleted remote branches:
```bash
git branch -vv | grep ': gone]'
```

6. To delete all local branches that track deleted remote branches:
```bash
git branch -vv | grep ': gone]' | awk '{print $1}' | xargs git branch -d
```