package balances

type Opts uint8

const DEFAULT = 0

const (
	VIRTUAL Opts = 1 << iota
	PASSIVE
	SYNCABLE
)

func (o Opts) IsVirtual() bool {
	return (o & VIRTUAL) != 0
}

func (o Opts) IsPassive() bool {
	return (o & PASSIVE) != 0
}

func (o Opts) IsActive() bool {
	return (o & PASSIVE) == 0
}

func (o Opts) IsSyncable() bool {
	return (o & SYNCABLE) != 0
}
