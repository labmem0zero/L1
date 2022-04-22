package points

type Point struct {
	x,y float64
}

func NewPoint(x,y float64)Point{
	pt:=Point{x,y}
	return pt
}

func (p *Point)SetXY(x,y float64){
	p.x,p.y=x,y
}

func (p *Point)GetXY()(x,y float64){
	return p.x,p.y
}

