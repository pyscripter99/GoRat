package output

import "go-rat/utils/types"

func OutputText(data string) (types.Output, error) {
	var output types.Output
	output.Element = "text"
	output.Data = data
	return output, nil
}
