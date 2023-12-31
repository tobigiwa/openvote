package chaincode

import "time"

// BioData is the data  of any person that can or cannot vote. Biodata
// is expected to be extended as demand requires.
type BioData struct {
	//DID Digital Identifier
	DID         string
	NationalID  string
	FirstName   string
	MiddleName  string
	LastName    string
	OtherName   string
	Sex         string
	BirthDay    int
	BirtMonth   int
	BirthYear   int
	Gender      string
	Citizenship string
}

// Age returns calculated age of person
func (b *BioData) Age() int {
	now := time.Now()
	age := now.Year() - b.BirthYear
	birthMonth := time.Month(b.BirtMonth)

	if now.Month() < birthMonth || (now.Month() == birthMonth && now.Day() < b.BirthDay) {
		age--
	}
	return age
}

func (b *BioData) EligibilityToVote() bool {
	if b.Age() < 18 {
		return false
	}
	return true
}

// Voter struct. Voter
// is expected to be extended as demand requires.
type Voter struct {
	Biodata             BioData
	VotedCandidateID    Candidate
	VoterPoliticalParty PoliticalParty
}

func (v Voter) HasBioData() BioData {
	return v.Biodata
}

// Candidate struct. Candidate
// is expected to be extended as demand requires.
type Candidate struct {
	BioData                 BioData
	CandidateID             string
	CandidatePoliticalParty PoliticalParty
	Votes                   uint64
}

func (c Candidate) HasBioData() BioData {
	return c.BioData
}

// PoliticalParty struct. PoliticalParty
// is expected to be extended as demand requires.
type PoliticalParty struct {
	PartyName        string
	PartyID          string
	PartyAbbrevation string
}

// Role is the contested position.
type Role string

// Election struct. Election
// is expected to be extended as demand requires.
type Election struct {
	ElectionYear int
	Position     string
	Contestants  []Candidate
	Winner       map[Role]Candidate
}


