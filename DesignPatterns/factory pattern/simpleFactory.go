package main

type Biryani interface {
	Cook() string
}

type chickenBiryani struct{}

func (c *chickenBiryani) Cook() string {
	return "Cooking Chicken Biryani"
}

type muttonBiryani struct{}

func (m *muttonBiryani) Cook() string {
	return "Cooking Mutton Biryani"
}

type vegBiryani struct{}

func (v *vegBiryani) Cook() string {
	return "Cooking Veg Biryani"
}

type BiryaniFactory struct{}

func (f *BiryaniFactory) GetBiryani(biryaniType string) Biryani {
	switch biryaniType {
	case "chicken":
		return &chickenBiryani{}
	case "mutton":
		return &muttonBiryani{}
	case "veg":
		return &vegBiryani{}
	default:
		return nil
	}
}

func SimpleFactory() {
	factory := &BiryaniFactory{}

	chickenBiryani := factory.GetBiryani("chicken")
	println(chickenBiryani.Cook())

	muttonBiryani := factory.GetBiryani("mutton")
	println(muttonBiryani.Cook())

}
