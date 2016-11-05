package keyboards

// Keyboard type regroup informations about keyboard on http://elitekeyboards.com
type Keyboard struct {
	Name      string
	Price     string
	Model     string
	Image     string
	Available bool
}

// Keyboards TODO: doc
type Keyboards map[string]Keyboard
