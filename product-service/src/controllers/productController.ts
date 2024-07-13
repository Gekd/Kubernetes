import { Request, Response } from 'express';

interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
}

let products: Product[] = [
  { id: 1, name: 'Product A', description: 'Description A', price: 100 },
  { id: 2, name: 'Product B', description: 'Description B', price: 150 }
];

export const getAllProducts = (req: Request, res: Response) => {
  res.json(products);
};

export const getProductById = (req: Request, res: Response) => {
  const id = parseInt(req.params.id);
  const product = products.find(p => p.id === id);
  if (product) {
    res.json(product);
  } else {
    res.status(404).send('Product not found');
  }
};

export const createProduct = (req: Request, res: Response) => {
  const newProduct: Product = req.body;
  newProduct.id = products.length + 1;
  products.push(newProduct);
  res.status(201).json(newProduct);
};

export const updateProduct = (req: Request, res: Response) => {
  const id = parseInt(req.params.id);
  const productIndex = products.findIndex(p => p.id === id);
  if (productIndex !== -1) {
    const updatedProduct: Product = { ...products[productIndex], ...req.body };
    products[productIndex] = updatedProduct;
    res.json(updatedProduct);
  } else {
    res.status(404).send('Product not found');
  }
};

export const deleteProduct = (req: Request, res: Response) => {
  const id = parseInt(req.params.id);
  const productIndex = products.findIndex(p => p.id === id);
  if (productIndex !== -1) {
    products.splice(productIndex, 1);
    res.status(204).send();
  } else {
    res.status(404).send('Product not found');
  }
};
