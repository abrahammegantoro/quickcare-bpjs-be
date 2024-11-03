package models

type ReferralType string

const (
	NewReferral       ReferralType = "New"
	ReferralExtension ReferralType = "Extension"
)

type Referral struct {
	Model
	Type                 ReferralType `bson:"type"`
	MedicalRecordFile    string       `bson:"medical_record_file"`
	VideoURL             *string      `bson:"video_url,omitempty"`
	PreviousReferralFile *string      `bson:"previous_referral_file,omitempty"`
	Explanation          *string      `bson:"explanation,omitempty"`
}
