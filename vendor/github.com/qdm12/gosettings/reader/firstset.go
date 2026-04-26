package reader

// FirstKeySet returns the first key set from a list of keys
// and from the given sources in order.
// It returns an empty string if none of the keys are set in
// any of the sources.
func (r *Reader) FirstKeySet(keys ...string) (firstKeySet string) {
	for _, key := range keys {
		for _, source := range r.sources {
			_, set := source.Get(key)
			if set {
				return key
			}
		}
	}
	return ""
}
