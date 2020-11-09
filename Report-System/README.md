# Report System
These commands are **not** standalone. Add all the commands if you wish to use them.
These CCs allow you to create a report system with the ability for users to request cancellation/nullification of their reports and add some functionalities for staff utilizing reactions.
All neccessary informations are composed in an embed which gets edited accordingly.

## Features
* Logging channel where report messages are sent into
* Notifying users about the current state of their report
* Enabling users to request cancellation of their report, in case of a mistake
* Utilization of reactions as menu options
* Editing the report message appropiately


## Setting Up

### Preface
Make for each custom command file a separate custom command, preferrably in the same category to keep them neat and organized. Please make sure to follow each step precisely and to use the correct trigger and trigger type, aswell.

#### These are the neccessary steps:
1. Disable the native report command, found here: `Control Panel > Tools & Utilities > Moderation`
    * I also recommend to create a command override disabling this command aswell
2. Copy the Channel-ID of the channel where you want your reports being logged into
    * paste it in the configuration area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$reportLog`
4. Copy the Channel-ID of the channel where you want to notify your members about the current state of their report
    *  paste it in the configuration area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$reportDiscussion`
5. Copy the Role-IDs of the roles which you consider administrators
    * paste them one by one in the configuration area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$adminRoles`
    * Separate the IDs with spaces
6. Copy the Role-IDs of the roles which you consider moderators
    * Copy the Role-IDs of the admin roles aswell
    * Paste them separated by spaces in the configuration area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$modRoles`
7. Run the case sensitive command `-ru dbSetup`
    * This command is restricted to admins only!
8. Done! YAGPDB.xyz will now take care of the rest and confirms setting up with an appropiate response

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
Each state has its own colour, for once to make it easier on the eyes and also to make it easier for you and your staff team recognizing in what state each report is.
* Pending moderator, not reviewed yet ![#FF00FF](https://via.placeholder.com/15/FF00FF/000000?text=+)
* Pending cancellation request ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+)
* Under investigation ![#FFFF00](https://via.placeholder.com/15/FFFF00/000000?text=+)
* Information requested ![#0000FF](https://via.placeholder.com/15/0000FF/000000?text=+)
* Report resolved (i.e. cancellation accepted, dismissal, action on reported user taken, and similar) ![#00FF00](https://via.placeholder.com/15/00FF00/000000?text=+)
* Cancellation request denied (defaults then, but with moderator) ![#FF0000](https://via.placeholder.com/15/FF0000/000000?text=+)

#### Default Reaction Inferface
![Default Interface Image](https://media.discordapp.net/attachments/767771719720632350/775133694264213523/unknown.png)

#### Reaction Inferface With Pending Cancellation Request
![Cancellation Inferface Image](https://media.discordapp.net/attachments/767771719720632350/775140298690134026/unknown.png)



## The Commands
### reactionHandler.go
[reactionHandler.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/reactionHandler.go) is the reaction listener and manages the editing of the report messages and user notifications. It is very fast. Like, *really* fast. If you react too eager on the message, it might break. Stay calm, nothing can happen. The report is safe as soon as the command is executed.

### customReport.go
[customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) manages the logging of the channel reported in, sends the reporting user a DM on how to cancel their report, sends the report to the according channel and primes this message with a guide and reactions.
This command tends to be rather slow, due to the massive priming of database entries and loads of messages and variables required. Compared to other commands with the same amount of DB and messaging, it still is on the faster end.

### cancelReport.go
[cancelReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/cancelReport.go) manages the cancellation requests. It still utilizes some database functions, but less than `customReport.go`. It primes the report message with the note that this report has a pending cancellation request and initializes another reaction menu, as well as a guide for that.
Should this command break, don't panic. You can still dismiss the report and choose "no action". This is just a functionality to prevent users getting warnings for false reports too easily.

## Small Disclaimer And Information
These commands are there to use "as-is" (after you did the configuration as described above), therefore they are not meant to be altered. You can still alter them (you should know what you are doing, though). In that case, however, I am not responsible for *any* bug or malfunction. I can still help you, if needed. For that, open an issue here or ping me in the [YAGPDB Support Server](https://discord.gg/5uVyq2E). Please allow me up to 1 day to answer, as I have school, like most other people, too. Oh, and timezones are also a thing.