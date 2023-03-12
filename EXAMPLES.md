### Examples

A couple of examples making use of the `upl` script.  
Modify `url` and `key` in the script and place it in a directory in your `PATH`

#### Upload a file

- `$ upl your.file`
- `$ upl your.file pfn` (preserves the filname)

#### Upload some programs output

Pipe some output to our script. `curl` reads and uploads data from stdin.

- `$ echo "some output" | upl`

#### Sway/Wayland: Upload screenshot

Uses `grim` and `slurp` to take a screenshot from a selected area.  
Let's you enrich the screenshot via `swappy` and then uploads it after closing `swappy`.

- `$ grim -g "$(slurp)" - | swappy -f - -o - | upl`

#### Upload clipboard content

Use your favorite clipboard tool and pipe the output to our script.

- `$ wl-paste | upl`
- `$ xclip -o | upl`
- `$ xsel -o | upl`
