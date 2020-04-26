// This package converst DNA strands to RNA strands
package strand

// ToRNA will return a RNA strand
func ToRNA(dna string) string {
	data := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}

	rna := ""
	for i := range dna {
		rna += data[string(dna[i])]
	}
	return rna
}
