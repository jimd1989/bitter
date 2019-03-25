// Access token, html constants, etc. A place to dump ugly strings that don't respect the 80 chars rule.

package main

import (
    "html/template"
)

// Your OAuth2 Bearer Token goes here. It can be generated once then used as a constant. This should not be exposed to the end-user in any way.
var TOKEN = "Bearer TOKENHERE"

// The URL for the tweet GET request
var T_URL = "https://api.twitter.com/1.1/statuses/show.json?&tweet_mode=extended&id="

// HTML that will be written to the top of every page
var HTMLHead = `
<html>
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0 maximum-scale=10.0 user-scalable=yes">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <style>
      body{padding:1em;padding-top:3em;max-width:52em;margin:0 auto;background:#e6ecf0;}
      #tweet{padding:1em;background:white;width=100%;}
    </style>
    <title>Bitter</title>
  </head>
  <body>
    `

// Closing HTML tags    
var HTMLFoot = `
  </body>
</html>
`

// The form that is displayed when bitter is provided no arguments
var SearchHTML = HTMLHead +
`
<div id="tweet">
  <h1>Bitter</h1>
  <p>Display a single tweet of the form "https://twitter.com/username/status/id" in the format you wish:</p>
  <form action="/" method="get">
    <select name="format">
      <option value="html" selected>HTML</option>
      <option value="json">JSON</option>
      <option value="">Text</option>
    </select>
    <input type="text" name="url" placeholder="Tweet URL">
    <button type="submit">Read</button>
  <form>
</div>
` +
HTMLFoot

// The tweet HTML view template
var TweetTemplate,  _ = template.New("Tweet").Parse(
    HTMLHead +
    `
    <div id="tweet">
      <p>{{.CreatedAt}}</p>
      <p><a href="https://twitter.com/{{.User.ScreenName}}">{{.User.ScreenName}}</a></p>
      <p>{{.FullText}}</p>
      <ul>
        {{range .Entities.Media }}
        <li><a href="{{.MediaURLHttps}}">{{.Type}}</a></li>
        {{end}}
      </ul>
    </div>
    ` +
    HTMLFoot)
