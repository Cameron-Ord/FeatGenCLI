package calculation

func Prepare_Data(data InputData) ([]string, error) {
	var allAllocatedFeats []string
	var Feats []map[string]string
	var err error
	Feats = Unload_Json()
	err = Run_Calculation(Feats, data, &allAllocatedFeats)
	if err != nil {
		return nil, err
	}

	return allAllocatedFeats, nil
}
