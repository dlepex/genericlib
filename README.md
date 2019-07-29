# genericlib

See https://github.com/dlepex/typeinst

## Overview

genericlib provides the following generic packages:

1. **slice**
   - basic operations: IndexOf, Contains, Delete, DeleteAt, Copy, Pop, Reverse
   - predicative operations: Filter, FilterMut, FilterTo, FindIndex, Exists, All
   - map, reduce
   - iterating slice by chunks
2. **conv**: *conversions between data sructures*
	- map keys/values to slice
	- slice to map, to set
	- set to slice
3. **set**:	*handy wrapper type for* ```map[E]struct{}```
