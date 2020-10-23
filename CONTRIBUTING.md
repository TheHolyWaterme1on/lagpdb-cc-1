# Contributing
This is a general list of personal standards I use for my code. However, I don't expect contributors to use the same standards as I personnaly write my CCs with.

## How To
First, you will need to follow the instructions **below**. Then, open a pull request!

## What you need to do
There are four things I expect you to do:
* add a leading comment at the top of your file
* add your CC in the appropiate section, named `ccName.go` where ccName is your CC name
* If your command requires any configuration (variables, etc.), explain that in the syntax provided below
* If it is a system of commands, add a new folder **and** add a `README.md` file in it

### Leading comment syntax
The following syntax is recommend for comments:
```go
{{/*
    <Comment>
*/}}
```
Note the indent and the space right after the comment. The comment should be like this:
```go
<Description> 

<Usage>

Recommended trigger: <Trigger type> trigger with trigger `<trigger>`.
```

If you have any doubts, feel free to ask me or look at the syntax in the existing files.

An example of a complete leading comment would be:
```go
{{/*
    This command sends the message "hello world" in the channel. 
    
    Usage: `-helloworld`

    Recommended trigger: Command trigger with trigger `helloworld`.

    Created by: yourName#0000 (optional: User ID aswell)
*/}}
```

### Configuration
If your command requires configuration, it should be done through appropiately named variables at the TOP of the file, right after the leading comment (as described above). For example:

```go
// Leading comment
{{/* CONFIGURATION AREA START */}}
{{ $location := "The moon" }} {{/* Appropriate description of what this variable does */}}
<Other configuration values>
{{/* CONFIGURATION AREA END */}}

// The code
```

That's all you need to do! Thanks for contributing!

## My personal guidelines for CCs
*Note: I do not require any of these items in your pull request, only the things stated above*

This is the guideline I use for the CCs I create. I may refactor your code when I'm doing a routine cleanup, so to avoid any confusion over what I am doing:
* `print` or `printf` for string concatenation rather than `joinStr` - joinStrt should generally be only used for joining a string slice.
* spaces after `{{` and `}}` (only one, no spaces after `(` and `)`)
* Printing the error message like a native error message