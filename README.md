# mirrorbot

> [!NOTE]
> Still in early development, the logic can update, change or be unstable

A lightweight GitHub Action to automatically sync issues from a GitHub mirror repository to a GitLab source repository.

## How to add the workflow

Create a file named `.github/workflows/mirror.yml` in your repository:

```yaml
name: Mirror to GitLab

on:
  issues:
    types: [opened]

jobs:
  mirror:
    runs-on: ubuntu-latest
    steps:
      - name: Sync Issue to GitLab
        uses: denis1836/mirrorbot@v1
        with:
          gitlab_token: ${{ secrets.GITLAB_TOKEN }}
          gitlab_project_id: "<project id>"
```

## License

[MIT](./LICENSE.md)
