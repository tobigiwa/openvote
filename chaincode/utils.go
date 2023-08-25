package chaincode

// HasBioDataInterface constraint for type that as Biodata type
type HasBioDataInterface interface {
	//HasBioData returns BioData type
	HasBioData() BioData
}

// Key returns voter.Biodata.NationalID if voter.Biodata.DID is empty.
func Key(entity HasBioDataInterface) string {
	if entity.HasBioData().DID != "" {
		return entity.HasBioData().NationalID
	}
	return entity.HasBioData().DID
}
