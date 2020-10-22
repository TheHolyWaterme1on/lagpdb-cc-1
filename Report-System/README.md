# Report System
These commands are **not** standalone. Add all the commands if you wish to use them.
These CCs allow you to create a report system with the ability for users to request cancellation/nullification of their reports and add some functionalities for staff utilizing reactions.

## Functionality
* Set logging channel where the reports are being logged into
* Notify users about their actions being taken on their reports
    * Set "talk-to-staff-channel" where users are being notified and can talk to the moderators

## Setting Up
1. Copy the ID of your report Logging channel (The one where the native report feature logs those reports into)
2. Disable the native report command in your `control panel > Tools & Utilities > Moderation`
    * I also recommend making a command override which disables that command, too
3. Paste it in the config area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$reports`
4. Copy the ID of your report-discussion channel (if you don't have one, make one!)
5. Paste it in the config area of [customReport.go](https://github.com/Olde7325/lagpdb-cc/blob/main/Report-System/customReport.go) to `$reportDiscussion`
6. YAGPDB will take care that the other commands know what these channels are
7. Done!