package framework

type AddWorkerArgs struct {
	Address string
}
type AddWorkerReply struct {
	ID int
}


type UploadArgs struct {
	Dir string
	Dt DataType
}
type UploadReply struct {}


type AssignDataArgs struct {
	Dir string
	FileName []string
	Dt DataType
}
type AssignDataReply struct {}


type PreprocessArgs struct {
	Dt DataType
}
type PreprocessReply struct{}