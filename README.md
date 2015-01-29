# Prolific

A little tool for authoring many tracker stories.

Prolific converts files that look like:

```
As a user I can toast a bagel

When I insert a bagel into toaster and press the on button, I should get a toasted bagel

L: mvp, toasting

---

As a user I can set the desired color of my bagel

I should be able to manipulate a dial and choose one of:

- light
- medium
- dark

Pressing the on button gives me toast of the appropriate color.

L: mvp, toasting

---

As a user I can clean my bagel toaster

I should be able to pull out a tray and clean up the crumbs.

L: mvp, clean-up
```

into CSV files ready for import into Tracker.  The first line becomes the story title, subsequent lines become content for the description, and the comma separated list after L: gets turned into labels.

Nothing fancy here.  Just a CSV file that you manually import into tracker.

## Usage

#### `prolific template`  

Will generate a template `stories.prolific` file

#### `prolific "Author Name" "path/to/stories.prolific"`

Will emit a CSV version of the passed in prolific file.  You can use `>` to shovel this content into a file.  For example:

```
prolific "Onsi Fakhouri" stories.prolific > stories.csv
```

The author name is used to populate the requester field on the story.  Make sure to use quotes!

## Installation

To install from source, make sure you have the Go toolchain installed, then:
`go install github.com/onsi/prolific`

Or just download the OS X binary from the GitHub releases page.

## License
Prolific is MIT Licenses