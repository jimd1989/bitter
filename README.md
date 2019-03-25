I don't like Twitter; it's over-engineered and everybody there is really unpleasant. It would be easy enough to ignore if it weren't so essential. Somewhere along the line journalists stopped writing actual articles and started stringing tweets together instead. All roads lead to Twitter if you want to keep on top of current events.

I wrote bitter (the Kiwi pronunciation of "better") to make the process of reading a tweet you've been linked to as painless as possible. It's an API with a simple web frontend that returns tweets (and attached media) as HTML, JSON, or plaintext. It's a lot easier on your bandwidth and strips away all "engagement" options such as liking, retweeting, or even seeing who responded to the tweet. It gets you in and out of the Twitterverse with your sanity intact.

Of course the ultimate solution is to simply stop caring about the news, but I'm not quite there yet.

## Installation

You will need [Go](https://golang.org) to build bitter. You will also need a Twitter developer account and [Bearer Token](https://developer.twitter.com/en/docs/basics/authentication/overview/application-only.html), which you should set as the value of `TOKEN` in `html.go` before building.

The process:

+ edit `TOKEN` in `html.go`
+ `make`
+ `make install` (may have to be root)
+ `make uninstall` (to remove)

## Usage

bitter is invoked as:

    bitter port

where `port` is the port bitter listens for connections.

## API

An HTML frontend is available when bitter is called with no parameters. It should be self-explanatory:

    https://bitter-url

bitter accepts two GET parameters: `url=` and `format=`. `url=` is a tweet URL of the form:

    https://twitter.com/username/status/id

and `format=` is either "html" for HTML output or "json" for JSON output. If `format=` is blank, it will respond with the tweet in an easily-parsed plaintext format, which may be better for command-line scripts that have no JSON facilities installed.

All media is returned as a link, so readers don't have to dip into their bandwidth unless they want to. Images are directly linked to, while videos link back to Twitter itself. Users are encouraged to paste the video link in [mpv](https://mpv.io/) to avoid this.

## Demo

You can try bitter [here](https://dalrym.pl/bitter), but the Twitter API limits a given token to a certain number of requests per hour. I encourage you to run your own instance instead.
