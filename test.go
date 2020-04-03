package main

import "github.com/rafaelmilanibarbosa/slack-go"
import "fmt"

func main() {
    webhookUrl := "https://hooks.slack.com/services/foo/bar/baz"

    attachment1 := slack.Attachment {}
    attachment1.AddField(slack.Field { Title: "Author", Value: "Rafael Barbosa" }).AddField(slack.Field { Title: "Status", Value: "Completed" })
    attachment1.AddAction(slack.Action { Type: "button", Text: "Book flights 🛫", Url: "https://www.google.com/flights", Style: "primary" })
    attachment1.AddAction(slack.Action { Type: "button", Text: "Cancel", Url: "https://www.google.com/flights", Style: "danger" })
    payload := slack.Payload {
      Text: "Hello from <https://github.com/ashwanthkumar/slack-go-webhook|slack-go-webhook>, a Go-Lang library to send slack webhook messages.\n<https://golangschool.com/wp-content/uploads/golang-teach.jpg|golang-img>",
      Username: "robot",
      Channel: "#go-messages",
      IconEmoji: ":monkey_face:",
      Attachments: []slack.Attachment{attachment1},
    }
    err := slack.Send(webhookUrl, "", payload)
    if len(err) > 0 {
      fmt.Printf("error: %s\n", err)
    }
}
