# Description:
**wgg** (WizGitGrep) is a compiled program for Windows that allows you to search for commits in a Git repository using a search term and display a list of commits with associated URLs. The user can select a commit from the list to open the URL in their configured browser.

# Prerequisites:
**Git**: You must have Git installed and accessible from the command line.
**Configured Browser**: You need to have a web browser installed and configured in the .env file so that it can open URLs.
**.env File**: You need to create an .env file that contains the necessary configurations for wgg to work correctly.
# Configuring the .env File
Create a file named .env in the same directory as the wgg executable. This file should contain the following environment variables:

PROJECT_PATH=C:\path\to\your\repository
The path to the Git repository where the command will be executed (in your local machine)

WORSKPACE=your-workspace
The workspace is the part of the url in your repository service where the code is uploaded, this variable is needed to construct the URL to open the browser

PROJECT=your-project
The project is also part of the url but is the current repository of git for your providers

BROWSER_PATH=C:\path\to\your\browser.exe
Browser path is the full path to the executable of the browser you want to use to open the URLs (chrome, edge, firefox, etc)

# Run and usage:
Just run in the terminal where the wgg executable is located: 
```bash
wgg
```
The program will prompt you for a search term:

Search: <keyword>

The program will search for all commits containing the search term in the Git history.

The program will display a list of commits with associated labels (letters) and their URLs, for example:
```bash
URL (A): https://github.com/your_username/your_project/a1b2c3d4e5
commit a1b2c3d4e5f67890123456789abcdef12345678
Author: Merlin Mage <merlin.mage@wizards.com>
Date:   Mon Sep 20 12:34:56 2023 -0400
    Fix bug in login feature
Select a Commit: Enter the letter of the commit you want to open in your configured browser:
```
At the end will prompt you to select a commit using the letter described above in the URL (Letter)
```bash 
Open commit: A
```

The program will open the browser to the URL corresponding to the selected commit.

# Example Configuration:
.env File:
PROJECT_PATH=C:\Users\your_username\git_project
WORSKPACE=github.com\Merlin-Apps
PROJECT=wizgitgrep
BROWSER_PATH=C:\Program Files\BraveSoftware\Brave-Browser\Application\brave.exe

# Notes:
If any of the required environment variables are not defined in the .env file, the program will display an error message.

The browser configured in BROWSER_PATH must be installed and accessible from the specified path.

Error loading the .env file: If the .env file fails to load, an error message will be displayed.

Error executing the git log command: If the Git command fails, an error message will be printed with the corresponding error.

Browser not found: If the configured browser cannot be executed, an error message will be shown, and the program will continue.

