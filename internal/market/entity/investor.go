package entity

type Investor struct {
	ID string
	Name string
	AssetPosition []*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	AssetsID string
	Shares int
}

func NewInvestor(id string) *Investor{
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	AssetPosition := i.GetAssetPosition(assetID)
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition{
	
}