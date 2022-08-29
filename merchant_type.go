package alipay

type ContactInfo struct {
	Name   string `json:"name,omitempty"`   // 联系人姓名
	Mobile string `json:"mobile,omitempty"` // 联系人手机
	Email  string `json:"email,omitempty"`  // 联系人邮箱
}

type DefaultSettleRule struct {
	DefaultSettleType   string `json:"default_settle_type"`   //默认结算类型，可选值有bankCard/alipayAccount。bankCard表示结算到银行卡；alipayAccount表示结算到支付宝账号
	DefaultSettleTarget string `json:"default_settle_target"` //默认结算目标。当默认结算类型为bankCard时填写银行卡卡号，其值需在进件填写的结算银行卡范围内；当默认结算类型为alipayAccount时填写支付宝账号登录号，其值需在进件填写的结算支付宝账号范围内。
}

type MerchantSimpleCreate struct {
	AppAuthToken           string            `json:"-"`                                   // 可选
	BindingAlipayLogonId   string            `json:"binding_alipay_logon_id"`             //签约支付宝账户
	AliasName              string            `json:"alias_name,omitempty"`                //商户别名
	ExternalId             string            `json:"external_id"`                         //商户编号，由机构定义，需要保证在机构下唯一
	ContactInfo            ContactInfo       `json:"contact_infos,omitempty"`             // 联系人信息
	Mcc                    string            `json:"mcc"`                                 //商户类别码mcc
	DefaultSettleRule      DefaultSettleRule `json:"default_settle_rule"`                 //默认结算规则
	AlipayLogonId          string            `json:"alipay_logon_id"`                     //结算支付宝账号，结算账号使用支付宝账号时必填。本字段要求与商户名称name同名，且是实名认证支付宝账户(个体工商户可以与name或cert_name相同)
	Service                []string          `json:"service"`                             //商户签约服务，可选值有：当面付、app支付、wap支付、电脑支付、线上资金预授权、新当面资金授权、商户代扣、小程序支付。其值会影响其他字段必填性，详见其他字段描述
	OutDoorImages          string            `json:"out_door_images,omitempty"`           //门头照，使用当面付服务时必填。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。
	InDoorImages           string            `json:"in_door_images,omitempty"`            //内景照，使用当面付服务时必填。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key
	AdditionalCertNo       string            `json:"additional_cert_no,omitempty"`        //补充证件号，与additional_cert_type+additional_cert_image同时提供。当商户类型为个人时，使用当面付收款有限额，补充这组证件信息可提额。目前仅允许个人类型商户传入本字段。
	AdditionalCertImage    string            `json:"additional_cert_image,omitempty"`     //补充证件图片，与additional_cert_no、additional_cert_type同时提供。当商户类型为个人时，使用当面付收款有限额，补充这组证件信息可提额。目前仅允许个人类型商户传入。其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。
	signTimeWithIsv        string            `json:"sign_time_with_isv,omitempty"`        //二级商户与服务商的签约时间。
	BizCards               SettleCardInfo    `json:"biz_cards,omitempty"`                 //结算银行卡信息，如果结算到支付宝账号，则不需要填写。本业务当前只允许传入一张结算卡。个人类型商户不允许结算到银行卡
	LicenseAuthLetterImage string            `json:"license_auth_letter_image,omitempty"` //授权函。当商户名与结算卡户名不一致（模板参考https://gw.alipayobjects.com/os/skylark-tools/public/files/d5fcbe7463d7159a0d362da417d157ed.docx），或涉及外籍法人（这种情况上传任意能证明身份的图片）时必填，其值为使用ant.merchant.expand.indirect.image.upload上传图片得到的一串oss key。
}

func (this MerchantSimpleCreate) APIName() string {
	return "ant.merchant.expand.indirect.zft.simplecreate"
}

func (this MerchantSimpleCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

type MerchantSimpleCreateRsp struct {
	Content struct {
		Code    Code   `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
		OrderId string `json:"order_id,omitempty"` // 申请单id
	} `json:"ant_merchant_expand_indirect_zft_simplecreate_response"`
	Sign string `json:"sign"`
}

func (this *MerchantSimpleCreateRsp) IsSuccess() bool {
	if this.Content.Code == CodeSuccess {
		return true
	}
	return false
}

type MerchantQuery struct {
	OrderId    string `json:"order_id"`              //申请单id。通过 ant.merchant.expand.indirect.zft.create(直付通二级商户创建)接口返回。
	ExternalId string `json:"external_id,omitempty"` //进件申请时的外部商户id，与order_id二选一必填
}

func (this MerchantQuery) APIName() string {
	return "ant.merchant.expand.indirect.zftorder.query"
}

func (this MerchantQuery) Params() map[string]string {
	var m = make(map[string]string)
	return m
}

type MerchantQueryRsp struct {
	Content struct {
		Code    Code            `json:"code"`
		Msg     string          `json:"msg"`
		SubCode string          `json:"sub_code"`
		SubMsg  string          `json:"sub_msg"`
		Orders  []MerchantOrder `json:"orders,omitempty"` // 直付通二级商户进件申请单信息
	} `json:"ant_merchant_expand_indirect_zftorder_query_response"`
	Sign string `json:"sign"`
}

func (this *MerchantQueryRsp) IsSuccess() bool {
	if this.Content.Code == CodeSuccess {
		return true
	}
	return false
}

type SettleCardInfo struct {
	AccountHolderName   string `json:"account_holder_name"`   //卡户名
	AccountNo           string `json:"account_no"`            //卡号
	AccountInstProvince string `json:"account_inst_province"` //开户行所在地-省
	AccountInstCity     string `json:"account_inst_city"`     //开户行所在地-市
	AccountBranchName   string `json:"account_branch_name"`   //开户支行名
	UsageType           string `json:"usage_type"`            //账号使用类型 对公-01  对私-02
	AccountType         string `json:"account_type"`          //卡类型   借记卡-DC 信用卡-CC
	AccountInstName     string `json:"account_inst_name"`     //银行名称
	AccountInstId       string `json:"account_inst_id"`       //银行名称
	BankCode            string `json:"bank_code"`             //联行号
}

type RoyaltyEntity struct {
	Type    string `json:"type"`           //分账接收方方类型。 userId：表示是支付宝账号对应的支付宝唯一用户号；loginName：表示是支付宝登录号
	Account string `json:"account"`        //分账接收方账号。 当分账方类型是userId时，本参数为用户的支付宝账号对应的支付宝唯一用户号，以2088开头的纯16位数字； 当分账方类型是loginName时，本参数为用户的支付宝登录号。
	Name    string `json:"name,omitempty"` /**分账接收方真实姓名。
	绑定分账关系时：
	当分账方类型是userId时，本参数可以不传，若上传则进行校验不上传不会校验。
	当分账方类型是loginName时，本参数必传。
	解绑分账关系时：作为请求参数可不填，分账关系查询时不作为返回结果返回
	**/
	Memo          string `json:"memo,omitempty"`          //分账关系描述
	LoginName     string `json:"loginName,omitempty"`     //作为查询返回结果：当前userId对应的支付宝登录号。当login_name与bind_login_name不相等时，表明该支付宝账户发生了登录号变更。
	BindLoginName string `json:"bindLoginName,omitempty"` //作为查询返回结果：分账收款方绑定时的支付宝登录号。分账关系绑定（alipay.trade.royalty.relation.bind）时，通过type为loginName绑定传入的支付宝登录号，若使用userId绑定则不返回
}

type RoyaltyBind struct {
	AppAuthToken string          `json:"-"`              // 可选
	ReceiverList []RoyaltyEntity `json:"receiver_list"`  //分账接收方列表，单次传入最多20个
	OutRequestNo string          `json:"out_request_no"` //外部请求号，由商家自定义。32个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复
}

func (this RoyaltyBind) APIName() string {
	return "alipay.trade.royalty.relation.bind"
}

func (this RoyaltyBind) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

type RoyaltyUnBind struct {
	AppAuthToken string          `json:"-"`              // 可选
	ReceiverList []RoyaltyEntity `json:"receiver_list"`  //分账接收方列表，单次传入最多20个
	OutRequestNo string          `json:"out_request_no"` //外部请求号，由商家自定义。32个字符以内，仅可包含字母、数字、下划线。需保证在商户端不重复
}

func (this RoyaltyUnBind) APIName() string {
	return "alipay.trade.royalty.relation.unbind"
}

func (this RoyaltyUnBind) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

type RoyaltyBindRsp struct {
	Content struct {
		Code       Code   `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		ResultCode string `json:"result_code,omitempty"` //SUCCESS：分账关系绑定成功； FAIL：分账关系绑定失败。
	} `json:"alipay_trade_royalty_relation_bind_response"`
	Sign string `json:"sign"`
}

type MerchantOrder struct {
	OrderId      string `json:"order_id"`      //申请单id
	ExternalId   string `json:"external_id"`   //外部商户id
	MerchantName string `json:"merchant_name"` //进件时填写的商户名称
	Status       string `json:"status"`        //申请总体状态。99:已完结;-1:失败;031:审核中
	ApplyTime    string `json:"apply_time"`    //申请单创建时间 2017-11-11 12:00:00
	FkAudit      string `json:"fk_audit"`      //风控审核状态。CREATE：已创建待审批、SKIP：跳过风控审批步骤、PASS：风控审核通过、REJECT：风控审批拒绝
	FkAuditMemo  string `json:"fk_audit_memo"` //风控审批备注，如有则返回
	KzAudit      string `json:"kz_audit"`      //客资审核状态。CREATE：已创建待审批、SKIP：跳过客资审批步骤、PASS：客资审核通过、REJECT：客资审批拒绝
	KzAuditMemo  string `json:"kz_audit_memo"` //客资审批备注，如有则返回
	SubConfirm   string `json:"sub_confirm"`   //二级商户确认状态。CREATE：已发起二级商户确认、SKIP：无需确认、FAIL：签约失败、NOT_CONFIRM：商户未确认、FINISH签约完成
	CardAliasNo  string `json:"card_alias_no"` //进件生成的卡编号，在发起结算时可以作为结算账号
	Smid         string `json:"smid"`          //二级商户id。当总体申请状态status为99时，smid才算进件完成
	ApplyType    string `json:"apply_type"`    //本申请单的请求类型。一般可选值包括ZHIFUTONG_CONSULT（直付通商户预校验）/ZHIFUTONG_CREATE（直付通商户创建）/ZHIFUTONG_MODIFY（直付通商户修改）
	AppPreAuth   string `json:"app_pre_auth"`  //是否开通线上预授权
	FacePreAuth  string `json:"face_pre_auth"` //是否开通线下预授权
	IsFaceLimit  string `json:"is_face_limit"` //判断个人当面付权限版本，返回TRUE时表示是标准版，返回FALSE表示受限版s
	Reason       string `json:"reason"`        //申请单处理失败时，通过此此段返回具体的失败理由；与kf_audit_memo和kz_audit_memo配合使用
}
