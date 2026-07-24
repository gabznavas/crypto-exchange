package main

import "time"

type Order struct {
	Size      float64 // 0.5 BTC
	Bid       bool    // se é compra
	Limit     *Limit  // Limit{Price: 65000.0, Order1, }
	Timestamp int64
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

type Limit struct {
	Price float64 // 65000.0

	// cada Order é de uma pessoa, por exemplo, mas todas partilham do mesmo price que pode ser de compra ou venda
	// Order{Size:0.5},Order{Size:0.2},Order{Size:0.5},
	Orders []*Order

	// imagino que total de volume aqui de order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

type Orderbook struct {
	// Eu vendo até X valor por isso
	// Do menor para o maior, quem cobra menos tem prioridade
	Asks []*Limit

	// Eu compro até X valor por isso
	// Do maior para o menor, quem paga mais, tem prioridade
	// a compra é feita aqui quando Bid >= que Ask
	Bids []*Limit
}

// exemplo:
// order e limit anda sempre junto então. Eu quero vender 0.5BTC a 65000.0.
// Então vou ter um limit dentro de ask de 65000 e uma order de 0.5BTC.
// Orderbook{
//   Asks: []*Limit{
//     &Limit{
//       Price: 65000.0,
//       Orders: []*Order{
//         &Order{ Size: 0.5, Limit: <ponteiro pro Limit acima> },
//       },
//       TotalVolume: 0.5,
//     },
//   },
// }
