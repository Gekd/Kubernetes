export interface Order {
    id: number;
    customerId: number;
    products: { productId: number; quantity: number }[];
    totalAmount: number;
    status: string;
  }
  
  export interface CreateOrder {
    customerId: number;
    products: { productId: number; quantity: number }[];
    totalAmount: number;
    status: string;
  }
  