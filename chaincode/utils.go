package chaincode

// Key returns voter.Biodata.NationalID if voter.Biodata.DID is empty.
func Key(voter Voter) string {
	if voter.Biodata.DID != "" {
		return voter.Biodata.NationalID
	}
	return voter.Biodata.DID
}
