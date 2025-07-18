// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.28.0
// source: accounter.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Type int32

const (
	Type_None    Type = 0
	Type_Income  Type = 1
	Type_Expense Type = 2
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "None",
		1: "Income",
		2: "Expense",
	}
	Type_value = map[string]int32{
		"None":    0,
		"Income":  1,
		"Expense": 2,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_accounter_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_accounter_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{0}
}

type Category int32

const (
	Category_Default       Category = 0
	Category_Game          Category = 1
	Category_Food          Category = 2
	Category_Travel        Category = 3
	Category_Education     Category = 4
	Category_Health        Category = 5
	Category_Shopping      Category = 6
	Category_Other         Category = 7
	Category_Transport     Category = 8
	Category_Entertainment Category = 9
	Category_Investment    Category = 10
	Category_Loan          Category = 11
	Category_Salary        Category = 12
	Category_OtherIncome   Category = 13
	Category_App           Category = 14
	Category_House         Category = 15
	Category_Utility       Category = 16
	Category_Gift          Category = 17
	Category_Snacks        Category = 18
)

// Enum value maps for Category.
var (
	Category_name = map[int32]string{
		0:  "Default",
		1:  "Game",
		2:  "Food",
		3:  "Travel",
		4:  "Education",
		5:  "Health",
		6:  "Shopping",
		7:  "Other",
		8:  "Transport",
		9:  "Entertainment",
		10: "Investment",
		11: "Loan",
		12: "Salary",
		13: "OtherIncome",
		14: "App",
		15: "House",
		16: "Utility",
		17: "Gift",
		18: "Snacks",
	}
	Category_value = map[string]int32{
		"Default":       0,
		"Game":          1,
		"Food":          2,
		"Travel":        3,
		"Education":     4,
		"Health":        5,
		"Shopping":      6,
		"Other":         7,
		"Transport":     8,
		"Entertainment": 9,
		"Investment":    10,
		"Loan":          11,
		"Salary":        12,
		"OtherIncome":   13,
		"App":           14,
		"House":         15,
		"Utility":       16,
		"Gift":          17,
		"Snacks":        18,
	}
)

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}

func (x Category) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Category) Descriptor() protoreflect.EnumDescriptor {
	return file_accounter_proto_enumTypes[1].Descriptor()
}

func (Category) Type() protoreflect.EnumType {
	return &file_accounter_proto_enumTypes[1]
}

func (x Category) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Category.Descriptor instead.
func (Category) EnumDescriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{1}
}

type AddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          Type                   `protobuf:"varint,5,opt,name=type,proto3,enum=accounter.v1.Type" json:"type,omitempty"`
	Category      Category               `protobuf:"varint,6,opt,name=category,proto3,enum=accounter.v1.Category" json:"category,omitempty"`
	Desc          string                 `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Amount        float64                `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Date          string                 `protobuf:"bytes,7,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	mi := &file_accounter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_None
}

func (x *AddRequest) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_Default
}

func (x *AddRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *AddRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AddRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type AddReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReply) Reset() {
	*x = AddReply{}
	mi := &file_accounter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReply) ProtoMessage() {}

func (x *AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReply.ProtoReflect.Descriptor instead.
func (*AddReply) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{1}
}

func (x *AddReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          Type                   `protobuf:"varint,1,opt,name=type,proto3,enum=accounter.v1.Type" json:"type,omitempty"`
	Category      Category               `protobuf:"varint,2,opt,name=category,proto3,enum=accounter.v1.Category" json:"category,omitempty"`
	StartDate     string                 `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       string                 `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Page          int32                  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	mi := &file_accounter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_None
}

func (x *ListRequest) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_Default
}

func (x *ListRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *ListRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *ListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          Type                   `protobuf:"varint,2,opt,name=type,proto3,enum=accounter.v1.Type" json:"type,omitempty"`
	Category      Category               `protobuf:"varint,3,opt,name=category,proto3,enum=accounter.v1.Category" json:"category,omitempty"`
	Desc          string                 `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Amount        float64                `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Date          string                 `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	mi := &file_accounter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{3}
}

func (x *Transaction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transaction) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_None
}

func (x *Transaction) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_Default
}

func (x *Transaction) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Transaction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ListReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Transactions  []*Transaction         `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	mi := &file_accounter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{4}
}

func (x *ListReply) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *ListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListReply) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListReply) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type StatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartDate     string                 `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       string                 `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	mi := &file_accounter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{5}
}

func (x *StatsRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *StatsRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type CategoryStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Category      Category               `protobuf:"varint,1,opt,name=category,proto3,enum=accounter.v1.Category" json:"category,omitempty"`
	CategoryName  string                 `protobuf:"bytes,2,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	Amount        float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Count         int32                  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryStats) Reset() {
	*x = CategoryStats{}
	mi := &file_accounter_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryStats) ProtoMessage() {}

func (x *CategoryStats) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryStats.ProtoReflect.Descriptor instead.
func (*CategoryStats) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{6}
}

func (x *CategoryStats) GetCategory() Category {
	if x != nil {
		return x.Category
	}
	return Category_Default
}

func (x *CategoryStats) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *CategoryStats) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CategoryStats) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type StatsReply struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TotalIncome       float64                `protobuf:"fixed64,1,opt,name=total_income,json=totalIncome,proto3" json:"total_income,omitempty"`
	TotalExpense      float64                `protobuf:"fixed64,2,opt,name=total_expense,json=totalExpense,proto3" json:"total_expense,omitempty"`
	Balance           float64                `protobuf:"fixed64,3,opt,name=balance,proto3" json:"balance,omitempty"`
	IncomeByCategory  []*CategoryStats       `protobuf:"bytes,4,rep,name=income_by_category,json=incomeByCategory,proto3" json:"income_by_category,omitempty"`
	ExpenseByCategory []*CategoryStats       `protobuf:"bytes,5,rep,name=expense_by_category,json=expenseByCategory,proto3" json:"expense_by_category,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *StatsReply) Reset() {
	*x = StatsReply{}
	mi := &file_accounter_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsReply) ProtoMessage() {}

func (x *StatsReply) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsReply.ProtoReflect.Descriptor instead.
func (*StatsReply) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{7}
}

func (x *StatsReply) GetTotalIncome() float64 {
	if x != nil {
		return x.TotalIncome
	}
	return 0
}

func (x *StatsReply) GetTotalExpense() float64 {
	if x != nil {
		return x.TotalExpense
	}
	return 0
}

func (x *StatsReply) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *StatsReply) GetIncomeByCategory() []*CategoryStats {
	if x != nil {
		return x.IncomeByCategory
	}
	return nil
}

func (x *StatsReply) GetExpenseByCategory() []*CategoryStats {
	if x != nil {
		return x.ExpenseByCategory
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_accounter_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteReply) Reset() {
	*x = DeleteReply{}
	mi := &file_accounter_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReply) ProtoMessage() {}

func (x *DeleteReply) ProtoReflect() protoreflect.Message {
	mi := &file_accounter_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReply.ProtoReflect.Descriptor instead.
func (*DeleteReply) Descriptor() ([]byte, []int) {
	return file_accounter_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_accounter_proto protoreflect.FileDescriptor

var file_accounter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x34, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xd4,
	0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x91, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d,
	0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x48, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x96,
	0x01, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x32, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x86, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6f,
	0x6d, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x62,
	0x79, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x11, 0x65,
	0x78, 0x70, 0x65, 0x6e, 0x73, 0x65, 0x42, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x27, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x29, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x65,
	0x6e, 0x73, 0x65, 0x10, 0x02, 0x2a, 0xfb, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x6f,
	0x64, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x10, 0x03, 0x12,
	0x0d, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x10, 0x09, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x6e, 0x10, 0x0b, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x10, 0x0d, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x70, 0x70, 0x10, 0x0e, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x10, 0x0f,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x10, 0x10, 0x12, 0x08, 0x0a,
	0x04, 0x47, 0x69, 0x66, 0x74, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x6e, 0x61, 0x63, 0x6b,
	0x73, 0x10, 0x12, 0x32, 0xee, 0x02, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x55, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x55, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x51, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x60, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x5f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounter_proto_rawDescOnce sync.Once
	file_accounter_proto_rawDescData = file_accounter_proto_rawDesc
)

func file_accounter_proto_rawDescGZIP() []byte {
	file_accounter_proto_rawDescOnce.Do(func() {
		file_accounter_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounter_proto_rawDescData)
	})
	return file_accounter_proto_rawDescData
}

var file_accounter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_accounter_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_accounter_proto_goTypes = []any{
	(Type)(0),             // 0: accounter.v1.Type
	(Category)(0),         // 1: accounter.v1.Category
	(*AddRequest)(nil),    // 2: accounter.v1.AddRequest
	(*AddReply)(nil),      // 3: accounter.v1.AddReply
	(*ListRequest)(nil),   // 4: accounter.v1.ListRequest
	(*Transaction)(nil),   // 5: accounter.v1.Transaction
	(*ListReply)(nil),     // 6: accounter.v1.ListReply
	(*StatsRequest)(nil),  // 7: accounter.v1.StatsRequest
	(*CategoryStats)(nil), // 8: accounter.v1.CategoryStats
	(*StatsReply)(nil),    // 9: accounter.v1.StatsReply
	(*DeleteRequest)(nil), // 10: accounter.v1.DeleteRequest
	(*DeleteReply)(nil),   // 11: accounter.v1.DeleteReply
}
var file_accounter_proto_depIdxs = []int32{
	0,  // 0: accounter.v1.AddRequest.type:type_name -> accounter.v1.Type
	1,  // 1: accounter.v1.AddRequest.category:type_name -> accounter.v1.Category
	0,  // 2: accounter.v1.ListRequest.type:type_name -> accounter.v1.Type
	1,  // 3: accounter.v1.ListRequest.category:type_name -> accounter.v1.Category
	0,  // 4: accounter.v1.Transaction.type:type_name -> accounter.v1.Type
	1,  // 5: accounter.v1.Transaction.category:type_name -> accounter.v1.Category
	5,  // 6: accounter.v1.ListReply.transactions:type_name -> accounter.v1.Transaction
	1,  // 7: accounter.v1.CategoryStats.category:type_name -> accounter.v1.Category
	8,  // 8: accounter.v1.StatsReply.income_by_category:type_name -> accounter.v1.CategoryStats
	8,  // 9: accounter.v1.StatsReply.expense_by_category:type_name -> accounter.v1.CategoryStats
	2,  // 10: accounter.v1.Accounter.Add:input_type -> accounter.v1.AddRequest
	4,  // 11: accounter.v1.Accounter.List:input_type -> accounter.v1.ListRequest
	7,  // 12: accounter.v1.Accounter.Stats:input_type -> accounter.v1.StatsRequest
	10, // 13: accounter.v1.Accounter.Delete:input_type -> accounter.v1.DeleteRequest
	3,  // 14: accounter.v1.Accounter.Add:output_type -> accounter.v1.AddReply
	6,  // 15: accounter.v1.Accounter.List:output_type -> accounter.v1.ListReply
	9,  // 16: accounter.v1.Accounter.Stats:output_type -> accounter.v1.StatsReply
	11, // 17: accounter.v1.Accounter.Delete:output_type -> accounter.v1.DeleteReply
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_accounter_proto_init() }
func file_accounter_proto_init() {
	if File_accounter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_accounter_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounter_proto_goTypes,
		DependencyIndexes: file_accounter_proto_depIdxs,
		EnumInfos:         file_accounter_proto_enumTypes,
		MessageInfos:      file_accounter_proto_msgTypes,
	}.Build()
	File_accounter_proto = out.File
	file_accounter_proto_rawDesc = nil
	file_accounter_proto_goTypes = nil
	file_accounter_proto_depIdxs = nil
}
