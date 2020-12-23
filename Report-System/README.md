# Report System
These commands are **not** standalone. Add all the commands if you wish to use them.
These CCs allow you to create a report system with the ability for users to request cancellation/nullification of their reports and add some functionalities for staff utilizing reactions.
All neccessary information is arranged in an embed which is edited accordingly.

# Table of Contents
<details>
<summary>Table of Contents</summary>

* [Features](#Features)
* [Setting Up](#Setting-Up)
    * [Preface](#Preface)
* [Usage](#Usage)
    * [Interface](#Interface)
        * [Reaction Menu](#Reaction-Menu)
        * [Colour Coding](#Colour-Coding)
        * [Default Reaction Inferface](#Default-Reaction-Inferface)
        * [Pending Cancellation Request](#Pending-Cancellation-Request)
</details>

## Features
* Logging channel where report messages are sent into
* Notifies users about the current state of their report
* Enables users to request cancellation of their latest report
* Shows report history of the reported user
* Utilizes reactions as menu options
* Report history of reported user
* Edits the report message appropiately


## Setting Up
### Preface
Make for each custom command file a separate custom command, preferrably in the same category to keep them neat and organized. Please make sure to follow each step precisely and to use the correct trigger and trigger type, as well.

#### These are the neccessary steps:
1. Disable the native report command, found here: `Control Panel > Tools & Utilities > Moderation`
    * I also recommend to create a command override disabling this command aswell
2. Copy the channel ID of the channel where you want your reports being logged into
    * paste it in the configuration area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$reportLog`
3. Copy the channel ID of the channel where you want to notify your members about the current state of their report
    *  paste it in the configuration area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$reportDiscussion`
4. Make sure to 
4. Run the case sensitive command `-ru dbSetup`
    * This command is restricted to admins only!
5. Done! YAGPDB.xyz will now take care of the rest and confirms setting up with an appropiate response.  
**Note:** Make sure to change `-` in both RegEx triggers to match YAGPDB's prefix in your server!

## Usage
### Commands
`-ru <User:Mention/ID> <Reason:Text>` - Sends the report. 

`-cr <MessageID:Text> <Key:Text> <Reason:Text>` - Requests cancellation of the report with that ID in connotation of that key. Only works for the latest report.

### Interface
#### Reaction Menu
* ❌ - Dismisses a report, you will be then prompted with the following;
    * ❗ - Warns the reporting user for a false report
    * 👌 - Closes report without any further action being taken
* ⚠️ - Requests further information, either regarding the report or the cancellation request
* 🛡️ - Starts Investigation of the reported issue
* 👍 - Resolves reported issue
* ✅ - Accepts cancellation request and closes report
* 🚫 - Denies cancellation request and goes back to the default report reaction menu

***
Once a report is closed, YAGPDB.xyz will add a white flag (🏳️) as reaction to signalize a closed report.

#### Colour Coding
Each state has its own colour, for one to make it easier on the eyes and also to make it easier for you and your staff team recognizing in what state each report is.

* ![#808080](https://via.placeholder.com/15/808080/000000?text=+) Pending moderator review
* ![#FF00FF](https://via.placeholder.com/15/FF00FF/000000?text=+) Pending cancellation request 
* ![#FFFF00](https://via.placeholder.com/15/FFFF00/000000?text=+) Under investigation 
* ![#0000FF](https://via.placeholder.com/15/0000FF/000000?text=+) Information requested
* ![#00FF00](https://via.placeholder.com/15/00FF00/000000?text=+) Report resolved 
* ![#FF0000](https://via.placeholder.com/15/FF0000/000000?text=+) Cancellation request denied


#### Default Reaction Inferface
![Default Interface Image](https://cdn.discordapp.com/attachments/767771719720632350/787880054238740530/unknown.png)

#### Pending Cancellation Request
![Cancellation Inferface Image](https://cdn.discordapp.com/attachments/767771719720632350/787880387141304350/unknown.png)

***Note:*** Upon the first report there will not be any report history. The images are purely examples and do not represent the reality.  
~~I was too lazy to reset my database, also I wanted to show how the report history is going to look like~~.