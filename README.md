# Git car for the [Bullettrain](https://github.com/bullettrain-sh/bullettrain-go-core) shell prompt

## Features:

- Displaying only when needed
- Current branch name
- State of the branch

**Callword**: `git`

**Template variables**:

* `.Icon`: the car's icon
* `.Name`: the name of the branch
* `.StatusIcon`: the status icons

**Template colours**:

* `c`: the car's colour
* `cs`: the car symbol's colour
* `css`: the status icon symbol's colour


## Car options

| Environment variable                   | Description                                                    | Default value                                                                        |
|:---------------------------------------|:---------------------------------------------------------------|:-------------------------------------------------------------------------------------|
| BULLETTRAIN_CAR_GIT_PAINT              | Colour override for the car't paint.                           | red:white                                                                            |
| BULLETTRAIN_CAR_GIT_TEMPLATE           | The car's template.                                            | `{{.Icon \| printf "%s " \| cs}}{{.Name \| c}}{{.StatusIcon \| printf " %s"\| csi}}` |
| BULLETTRAIN_CAR_GIT_SYMBOL_ICON        | Icon displayed on the car.                                     | ``                                                                                  |
| BULLETTRAIN_CAR_GIT_SYMBOL_PAINT       | Colour override for the car's symbol.                          | red:white                                                                            |
| BULLETTRAIN_CAR_GIT_DIRTY_ICON         | Icon displayed when there are changes.                         | `✘`                                                                                  |
| BULLETTRAIN_CAR_GIT_DIRTY_PAINT        | Colour override for the dirty symbol.                          | red:white                                                                            |
| BULLETTRAIN_CAR_GIT_CLEAN_ICON         | Icon displayed when there are no changes.                      | `✔`                                                                                  |
| BULLETTRAIN_CAR_GIT_CLEAN_PAINT        | Colour override for the clean symbol.                          | green:white                                                                          |
| BULLETTRAIN_CAR_GIT_SEPARATOR_PAINT    | Colour override for the car's right hand side separator paint. | Using default painting algorythm.                                                    |
| BULLETTRAIN_CAR_GIT_SEPARATOR_SYMBOL   | Override the car's right hand side separator symbol.           | Using global symbol.                                                                 |
| BULLETTRAIN_CAR_GIT_SEPARATOR_TEMPLATE | Defines the car separator's template.                          | Using global template.                                                               |

# Contribute

Even reporting your use case will greatly help us to figure out/improve
this product, so feel free to reach out in the Issues section.
