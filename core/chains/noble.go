package chains

func Noble() (int, error) {
	url := "https://rest.cosmos.directory/noble/cosmos/staking/v1beta1/validators?page.offset=1&pagination.limit=500&status=BOND_STATUS_BONDED"
	return fetchCosmosSDKNakaCoeff("noble", url)
}
