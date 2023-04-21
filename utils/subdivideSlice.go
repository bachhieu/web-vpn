package utils

func ChunkSlice(input []interface{}, chunkSize int) [][]interface{} {
	var chunks [][]interface{}

	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize

		if end > len(input) {
			end = len(input)
		}

		chunks = append(chunks, input[i:end])
	}

	return chunks
}
