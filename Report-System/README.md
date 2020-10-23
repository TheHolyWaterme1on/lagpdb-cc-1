# Report System
These commands are **not** standalone. Add all the commands if you wish to use them.
These CCs allow you to create a report system with the ability for users to request cancellation/nullification of their reports and add some functionalities for staff utilizing reactions.

## Functionality
* Set logging channel where the reports are being logged into
* Notify users about their actions being taken on their reports
    * Set "talk-to-staff-channel" where users are being notified and can talk to the moderators
* Edit report message to the current state (e.g. "under investigation")
* Use reactions as menu options

## Setting Up
1. Copy the ID of your report Logging channel (The one where the native report feature logs those reports into)
2. Disable the native report command in your `control panel > Tools & Utilities > Moderation`
    * I also recommend creating a command override which disables that command, too
3. Paste it in the config area of every command to `$reportLog`
4. Copy the ID of your report-discussion channel (if you don't have one, make one!)
5. Paste it in the config area of every command to `$reportDiscussion`
6. Done!

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