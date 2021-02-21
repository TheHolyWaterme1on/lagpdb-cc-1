# Contributing
First of all, thank you for considering a contribution! I appreciate any kind of contributions, be that a bugfix, a new command, a typo fix, or anything else.
However, please do take the time to read the following guidelines, as detailed below.

## Guidelines
There's a few things you have to do before opening a pull request:

### Leading Comment
Add a brief comment at the start of your custom command code explaining what your command does and any setup required. 
This should include the trigger & trigger type in addition to any restrictions recommended. 
This is also the place to put information about you, the author. This is the most complicated step, so we've including an example below.

Let's say this was your original code:

```go
{{sendMessage nil "Hello, World!"}}
```
Then this will become your code afterwards:
```go
{{/*
    This command sends "Hello World!" to the current channel.
    
    Usage: `-helloworld`

    Recommended trigger and trigger type: Command trigger with trigger `helloworld`

    Author: Luca Z. <https://github.com/l-zeuch>
    License: AGPL-3.0
*/}}

{{sendMessage nil "Hello, World!"}}
```
If you have a Discord account, please do not add it, as these may change at a more regular interval, especially the tags. It's rather hard to keep up with possible changes - best way is to include your Discord tag in your GitHub description, that way you can edit it whenever and don't have to open a new PR each time you change it : )

### Configuration Variables
Every variable that needs to be changed is to be extracted in a section at the very top of your code, right underneath the leading comment. Please enclose this configuration area with comments stating where it begins and where it ends, such as follows:

```go
{{/* Config area start */}}

{{$ROLE_ID := ######}} {{/* The role you want to check. */}}

{{/* Config area end */}}

{{if hasRoleID $ROLE_ID}}
	Yes, you have the role.
{{else}}
	No, you do not have the role.
{{end}}
```
We named our configuration variable `ROLE_ID`, in UPPER_SNAKE_CASE, as it makes it clear that it is a constant. Note that this is not required, just recommended.

You'll also notice that we added a description of the variable, what it does and what you expect the end user to give as data. The enclosing comments may be anything that expresses the start and end of such a configuration area.
When there is no configuration required, you of course don't need a configuration area. You can then state "no configuration required" in your leading comment, this however is not required.

### Formatting
Your code should be relatively readable. This of course does not mean it has to be indented beautifully or the like, just make sure people other than you can follow along easily.
Should it be too long, don't worry - I'll run a [minifier ](https://jo3-l.github.io/cc-minifier/) (made by [Joe L.](https://github.com/jo3-l)) on it, as that does a fairly good job getting it down to length.

## How do I request for my code to be added?
After you've followed these steps, you are now ready for a PR!  Take a look at [this article](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request9 if you aren't quite familiar with PRs or need a refresher. We have a PR template in place that we recommend you follow.

If you're PRing one single command, all you have to do is adding it in the appropiate folder. If that folder does not exist, feel free to add that as well.

PRing a system of commands requires a few more steps: Find an appropiate folder and create a new folder containing only your custom commands plus a README.md documenting your custom commands.
You can ignore the file extension, I'll take care of that!

Thanks for contributing!
