# mnemonic
golang library for naming things

# Mnemonic Encoding Wordlist as a Service

>There are only two hard things in Computer Science: cache invalidation and naming things.
>
>-- Phil Karlton


The nice folks at [mnx.io](http://mnx.io/) wrote a fun [blog post](http://mnx.io/blog/a-proper-server-naming-scheme/) about server naming schemes.

One of the pearls of wisdom is the recommendation to use Oren Tirosh’s [mnemonic encoding project](http://web.archive.org/web/20090918202746/http://tothink.com/mnemonic/wordlist.html) as a wordlist. 

GENIUS!

I’d never heard of that particular wordlist before and immediately downloaded it and built a [shell function](https://gist.github.com/alexlovelltroy/119c32a12f6aca28c3f3) to give me a few examples for codenames.

This very small golang library exposes the wordlist for your naming needs.

## Usage

Here's an example that makes the wordlist available as a webservice.

```
package main

import (
	"strconv"

	"github.com/alexlovelltroy/mnemonic"
	"github.com/gin-gonic/gin"
)

func getRandomWord(c *gin.Context) {
	wordcount := c.Params.ByName("count")
	intWordCount, err := strconv.Atoi(wordcount)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid word count.  Must be an integer."})
	}
	c.JSON(200, gin.H{
		"words": mnemonic.GetRandomWords(intWordCount),
	})
}

func main() {
	r := gin.Default()
	r.GET("/random/:count", getRandomWord)
	r.Run(":8080")
}

```

### Bonus

The Docker project, which is now called moby, has it's own method of generating names based on scientists.  I've always been a fan of that too.

https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go




