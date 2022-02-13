# Build instructions (windows)
Built with [Go](https://go.dev/).

```go
set GOOS=windows GOARCH=amd64; go build .
```

# Usage (windows)

Move `groupme-files.exe` into a suitable parent directory like `C:\Users\mmmveggies\Desktop`.

Navigate to that parent directory from Windows PowerShell or a Command Prompt window, e.g.
```
Microsoft Windows [Version 10.0.19042.1466]
(c) Microsoft Corporation. All rights reserved.

C:\Users\mmmveggies>cd Desktop

C:\Users\mmmveggies\Desktop>
```

Enter configuration setup on the first execute (or whenever you want to change your Access Token):
```
C:\Users\mmmveggies\Desktop>.\groupme-files.exe configure
Using config file: C:\Users\mmmveggies\AppData\Roaming\groupme-files\.groupme-files.env
v Access Token: █
```

Copy your API token from logging in to https://dev.groupme.com and clicking the `Access Token` modal in page header to reveal the 40-character alphanumeric key.

Paste this token into the command line -- by right clicking the command line window (ctrl+v does not seem to work) -- and press enter.
```
C:\Users\mmmveggies\Desktop>.\groupme-files.exe configure
Using config file: C:\Users\mmmveggies\AppData\Roaming\groupme-files\.groupme-files.env
Access Token: ****************************************
Access Token: ****************************************
Updated configuration in: C:\Users\mmmveggies\AppData\Roaming\groupme-files\.groupme-files.env
```

After configuration, the only other command is `download`.
Launch this to echo your configuration and pick a group's messages to export:
```
C:\Users\mmmveggies\Desktop>.\groupme-files.exe download
Using config file: C:\Users\mmmveggies\AppData\Roaming\groupme-files\.groupme-files.env
2022/02/13 11:51:03 token ****************************************
Use the arrow keys to navigate: ↓ ↑ → ←
? Pick a target group:
  > A (id: ********)
    B (id: ********)
    C (id: ********)
```

The next prompt asks for a good download location - the default is a 'groupme-downloads' subfolder in the same parent folder as the running application. Change if necessary but press enter to accept the defaults.
```
v Where should the files be downloaded to?: C:\Users\mmmveggies\Desktop/groupme-downloads
```

Enter the message age limits. The default values represent messages created only in the past month (note that older image attachments take longer to page back to)
```
What is the upper age limit? (YYYY-MM-DD): 2022-01-13
v What is the lower age limit? (YYYY-MM-DD): █022-02-13
```

You should see something similar to the following, which may go on for several minutes. Images are organized by message sender and named for the message's time stamp.
```
2022/02/13 12:17:54 start 2022-01-13 00:00:00 +0000 UTC
2022/02/13 12:17:54 end 2022-02-13 00:00:00 +0000 UTC
2022/02/13 12:17:55 Reading page starting at: 2022-02-11 15:39:49 +0000 UTC
2022/02/13 12:17:55 Downloaded: C:\Users\mmmveggies\Desktop/groupme-downloads/MyFriend/2022-02-11_00-12-44__00.jpeg
2022/02/13 12:17:55 Downloaded: C:\Users\mmmveggies\Desktop/groupme-downloads/MyFriend/2022-02-11_00-12-44__01.jpeg
2022/02/13 12:17:56 Downloaded: C:\Users\mmmveggies\Desktop/groupme-downloads/mmmveggies/2022-01-20_03-20-48__00.png
2022/02/13 12:17:56 Downloaded: C:\Users\mmmveggies\Desktop/groupme-downloads/mmmveggies/2022-01-20_01-48-27__00.png
```

If you run the program again, the file will not overwrite existing files - delete them yourself if you want to re-download.
```
2022/02/13 12:18:21 start 2022-01-13 00:00:00 +0000 UTC
2022/02/13 12:18:21 end 2022-02-13 00:00:00 +0000 UTC
2022/02/13 12:18:22 Reading page starting at: 2022-02-11 15:39:49 +0000 UTC
2022/02/13 12:18:22 File already exists: C:\Users\mmmveggies\Desktop/groupme-downloads/MyFriend/2022-02-11_00-12-44__00.jpeg
2022/02/13 12:18:22 File already exists: C:\Users\mmmveggies\Desktop/groupme-downloads/MyFriend/2022-02-11_00-12-44__01.jpeg
2022/02/13 12:18:22 File already exists: C:\Users\mmmveggies\Desktop/groupme-downloads/mmmveggies/2022-01-20_03-20-48__00.png
2022/02/13 12:18:22 File already exists: C:\Users\mmmveggies\Desktop/groupme-downloads/mmmveggies/2022-01-20_01-48-27__00.png
```