# Notifier Package for GoLang
This package provides easy way to send notifications using supported notifiers.

## Supported notifiers
As of now, only `Telegram` is supported, next one in line would be `Slack`.

# Notification object
We are using `Notification` structure to create new instance of notification which can later be used to send it to recipients using on of the notifiers.

## Creating notification object
To create new notification object, you need to call the following piece of code:
```go
notification := notifier.NewNotification(notifier.NotificationOptions{
    Image:   "URL to image",          // This field is optional
    Type:    notifier.Success,
    Title:   "Notification Title",
    Message: "Notification Message",
})
```

## Notification types
We are supporting 5 types of notifications out of the box, these are:
1. Info - ℹ️ - `notifier.Info`  
![Info Notification](http://github.com/leads-su/notifier/-/raw/main/docs/images/info.png?inline=true)
2. Success - ✅ - `notifier.Success`  
![Success Notification](http://github.com/leads-su/notifier/-/raw/main/docs/images/success.png?inline=true)
3. Warning - ⚠️ - `notifier.Warning`  
![Warning Notification](http://github.com/leads-su/notifier/-/raw/main/docs/images/warning.png?inline=true)
4. Error - ❌ - `notifier.Error`  
![Error Notification](http://github.com/leads-su/notifier/-/raw/main/docs/images/error.png?inline=true)
5. Unknown - ⁉️ - `notifier.Unknown`  
![Unknown Notification](http://github.com/leads-su/notifier/-/raw/main/docs/images/unknown.png?inline=true)

You can use them to style your notifications.

## Notification with image as a header
To create notification with the image as a header, you need to pass `Image` option to the notification.  
This in turn will create notification which will look similar to this one  
![Image Notification](http://github.com/leads-su/notifier/-/raw/main/docs/images/image.png?inline=true)

# Telegram Notifier
To create new instance of **Telegram** notifier, you have to obtain bot token and know IDs of users you want to send notifications to.

## Obtaining bot token
1. Start chat with the [@BotFather](https://t.me/botfather) bot
2. Send `/newbot` command to start creation process
3. Enter new name for the bot when asked - This is the pretty version of name for the bot
4. Enter new username for the bot - This is the username of bot, it must end with `_bot` suffix
5. Copy token string which is generated by the **@BotFather**

## Creating new instance of notifier
To create new instance of Telegram notifier, you simply have to call the following line of code.  
It will return new instance of Telegram notifier which exposes `DeliverNotification` method.
```go
telegram := notifier.NewTelegramNotifier("BOT_TOKEN")
```

## Sending notification to recipients
To send notification object created as described above, you need to call `DeliverNotification` method on your `telegram` notifier instance.  
To do so, you can use something like this:
```go
err := telegram.DeliverNotification(notification, []string{
    "12345678"
})
if err != nil {
    fmt.Printf("Failed to send notification: %s\n", err.Error())
}
```

Here, `notification` is an instance of `Notification` object, and `[]string{}` is a list of recipients for notification.
