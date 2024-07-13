export interface Payment {
    id: number;
    customerId: number;
    totalAmount: number;
    status: string;
  }
  
  export interface CreatePayment {
    customerId: number;
    products: { productId: number; quantity: number }[];
    totalAmount: number;
    status: string;
  }
  