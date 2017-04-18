## import

```html

import (
	"icons"
	icon "icons/icon.vgo"
	)

<div class="head">
	<div>ego:{{.head}}</div>

	<icon>
		vclass={icon-share-to}
		node={ id="slot1"}
		prpo={node---1}
	</icon>

	<div>
		{{range .parArr}}
	        <p>arr::: {{.}}</p>
		{{end}}
	</div>

</div>
```