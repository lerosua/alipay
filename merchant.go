package alipay

// LifeTradePay (生活代扣)统一收单交易支付接口 https://docs.open.alipay.com/api_1/alipay.trade.pay/
func (this *Client) LifeTradePay(param TradePay) (result *TradePayRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeRefund (生活代扣)统一收单交易退款接口 https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (this *Client) LifeTradeRefund(param TradeRefund) (result *TradeRefundRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeOrderSettle (生活代扣)统一收单交易结算接口 https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (this *Client) LifeTradeOrderSettle(param TradeOrderSettle) (result *TradeOrderSettleRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// TradeOrderSettle (生活代扣)交易分账查询接口 https://docs.open.alipay.com/api_1/alipay.trade.order.settle.query
func (this *Client) LifeTradeOrderSettleQuery(param TradeOrderSettleQuery) (result *TradeOrderSettleQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// ant.merchant.expand.indirect.zft.simplecreate(直付通二级商户免证照进件) [https://opendocs.alipay.com/open/02n330]
func (this *Client) MerchantSimpleCreate(param MerchantSimpleCreate) (result *MerchantSimpleCreateRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

//ant.merchant.expand.indirect.zft.consult(直付通商户创建预校验咨询)
func (this *Client) MerchantConsult(param MerchantSimpleCreate) (result *MerchantSimpleCreateRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

//ant.merchant.expand.indirect.zftorder.query(直付通商户入驻进度查询)
func (this *Client) MerchantQuery(param MerchantQuery) (result *MerchantQueryRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

//ant.merchant.expand.indirect.zft.delete(直付通二级商户作废)
func (this *Client) MerchantDelete(param MerchantDelete) (result *MerchantDeleteRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

// alipay.trade.royalty.relation.bind(分账关系绑定)
func (this *Client) RoyaltyBind(param RoyaltyBind) (result *RoyaltyBindRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

//alipay.trade.royalty.relation.unbind(分账关系解绑)
func (this *Client) RoyaltyUnBind(param RoyaltyUnBind) (result *RoyaltyBindRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}
