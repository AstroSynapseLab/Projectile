# Projectile 

Projectile is a simple CLI tool designed to manage environments. It streamlines the process of setting up and cloning projects, aiding in the quick deployment of custom environment builds.

With a flexible design, Projectile is capable of supporting different project architectures. It works with both monolith app projects and Service-Oriented Architecture (SOA) based projects. For a monolith project, a `.projectile` folder is required. For SOA projects, a `projectile` repository within the project organization is needed. 

Please note that, as of now, Projectile only supports projects hosted on GitHub.

> **Note:** Projectile is currently in early development.

## Installation

1. **Clone the Repository:**

    ```
    git clone https://github.com/AstroSynapseLab/Projectile ~/.projectile
    ```

2. **Create an Alias:**

    ```
    echo "alias projectile='~/.projectile/bin'" >> ~/.zshrc
    ```

    Be sure to replace `.zshrc` with your shell's configuration file if you're using a different shell.

### Login

To authenticate with GitHub, use the `projectile login` command. You'll need a GitHub Personal Access Token (PAT). If you don't have a PAT, you can follow [this guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) to create one.

```
projectile login
```

## Clone an Existing Project

To clone a project, use the `projectile clone` command. This command expects a GitHub URL that points to either a Projectile-compatible repository or an organization that contains a Projectile repository.

Here's an example:

```
projectile clone https://github.com/AstroSynapseAI
```

In this example, `https://github.com/AstroSynapseAI` is an organization that contains the `https://github.com/AstroSynapseAI/projectile` repository.