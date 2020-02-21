package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetFrame() BuildProcess
	Build() VehicleProduct
}
type ManufacturingDirector struct {
	builder BuildProcess
}

func (md *ManufacturingDirector) Construct() {
	md.builder.SetWheels().SetSeats().SetFrame()
}
func (md *ManufacturingDirector) SetBuilder(b BuildProcess) {
	md.builder = b
}

type VehicleProduct struct {
	Wheels int
	Seats  int
	Frame  string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetFrame() BuildProcess {
	c.v.Frame = "Car"
	return c
}
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetFrame() BuildProcess {
	b.v.Frame = "Bike"
	return b
}
func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}
