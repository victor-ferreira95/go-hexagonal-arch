package cli

import (
	"fmt"
	"go-hexagonal/application"
)

func Run(service application.ProductServiceInterface,
	action string,
	productId string,
	productName string,
	price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())

		return result, nil

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
		return result, nil

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
		return result, nil

	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s\nName %s\nPrice %.2f\nStatus %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	}

	return result, nil
}
