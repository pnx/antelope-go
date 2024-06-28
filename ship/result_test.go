package ship_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/shufflingpixels/antelope-go/abi"
	"github.com/shufflingpixels/antelope-go/chain"
	"github.com/shufflingpixels/antelope-go/internal/assert"
	"github.com/shufflingpixels/antelope-go/ship"
)

func TestStatusResponseEncode(t *testing.T) {
	req := ship.Result{
		StatusResult: &ship.GetStatusResultV0{
			Head: &ship.BlockPosition{
				BlockNum: 893,
				BlockID: [32]byte{
					0x52, 0x40, 0x67, 0x7a, 0x86, 0x2d, 0x5a, 0x4d, 0x99, 0x80, 0xfe, 0x60, 0xb0, 0x33, 0xa2, 0xda,
					0xf1, 0xb1, 0xac, 0x7a, 0xa8, 0x64, 0x7b, 0xac, 0x33, 0x06, 0xbb, 0x99, 0x83, 0x17, 0x1d, 0x75,
				},
			},
			LastIrreversible: &ship.BlockPosition{
				BlockNum: 857,
				BlockID: [32]byte{
					0xd1, 0xba, 0xa2, 0x3f, 0x59, 0xdc, 0xac, 0x4e, 0xb6, 0x9a, 0x98, 0x32, 0x93, 0x7f, 0x0c, 0x6c,
					0x8d, 0xdd, 0x88, 0x44, 0x42, 0x24, 0x45, 0x73, 0x8a, 0x39, 0x43, 0x64, 0xde, 0x70, 0x4a, 0x46,
				},
			},
			TraceBeginBlock:      1000,
			TraceEndBlock:        2000,
			ChainStateBeginBlock: 8000,
			ChainStateEndBlock:   9000,
		},
	}

	expected := []byte{
		0x00, 0x7d, 0x03, 0x00, 0x00, 0x52, 0x40, 0x67,
		0x7a, 0x86, 0x2d, 0x5a, 0x4d, 0x99, 0x80, 0xfe,
		0x60, 0xb0, 0x33, 0xa2, 0xda, 0xf1, 0xb1, 0xac,
		0x7a, 0xa8, 0x64, 0x7b, 0xac, 0x33, 0x06, 0xbb,
		0x99, 0x83, 0x17, 0x1d, 0x75, 0x59, 0x03, 0x00,
		0x00, 0xd1, 0xba, 0xa2, 0x3f, 0x59, 0xdc, 0xac,
		0x4e, 0xb6, 0x9a, 0x98, 0x32, 0x93, 0x7f, 0x0c,
		0x6c, 0x8d, 0xdd, 0x88, 0x44, 0x42, 0x24, 0x45,
		0x73, 0x8a, 0x39, 0x43, 0x64, 0xde, 0x70, 0x4a,
		0x46, 0xe8, 0x03, 0x00, 0x00, 0xd0, 0x07, 0x00,
		0x00, 0x40, 0x1f, 0x00, 0x00, 0x28, 0x23, 0x00,
		0x00,
	}

	buf := new(bytes.Buffer)
	enc := abi.NewEncoder(buf, abi.DefaultEncoderFunc)
	err := enc.Encode(req)
	assert.NoError(t, err)

	data, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Equal(t, data, expected)
}

func TestStatusResultDecode(t *testing.T) {
	data := []byte{
		0x00, 0x7d, 0x03, 0x00, 0x00, 0x52, 0x40, 0x67,
		0x7a, 0x86, 0x2d, 0x5a, 0x4d, 0x99, 0x80, 0xfe,
		0x60, 0xb0, 0x33, 0xa2, 0xda, 0xf1, 0xb1, 0xac,
		0x7a, 0xa8, 0x64, 0x7b, 0xac, 0x33, 0x06, 0xbb,
		0x99, 0x83, 0x17, 0x1d, 0x75, 0x59, 0x03, 0x00,
		0x00, 0xd1, 0xba, 0xa2, 0x3f, 0x59, 0xdc, 0xac,
		0x4e, 0xb6, 0x9a, 0x98, 0x32, 0x93, 0x7f, 0x0c,
		0x6c, 0x8d, 0xdd, 0x88, 0x44, 0x42, 0x24, 0x45,
		0x73, 0x8a, 0x39, 0x43, 0x64, 0xde, 0x70, 0x4a,
		0x46, 0xe8, 0x03, 0x00, 0x00, 0xd0, 0x07, 0x00,
		0x00, 0x40, 0x1f, 0x00, 0x00, 0x28, 0x23, 0x00,
		0x00,
	}

	req := ship.Result{}
	err := abi.NewDecoder(bytes.NewBuffer(data), abi.DefaultDecoderFunc).Decode(&req)
	assert.NoError(t, err)

	expected := ship.Result{
		StatusResult: &ship.GetStatusResultV0{
			Head: &ship.BlockPosition{
				BlockNum: 893,
				BlockID: [32]byte{
					0x52, 0x40, 0x67, 0x7a, 0x86, 0x2d, 0x5a, 0x4d, 0x99, 0x80, 0xfe, 0x60, 0xb0, 0x33, 0xa2, 0xda,
					0xf1, 0xb1, 0xac, 0x7a, 0xa8, 0x64, 0x7b, 0xac, 0x33, 0x06, 0xbb, 0x99, 0x83, 0x17, 0x1d, 0x75,
				},
			},
			LastIrreversible: &ship.BlockPosition{
				BlockNum: 857,
				BlockID: [32]byte{
					0xd1, 0xba, 0xa2, 0x3f, 0x59, 0xdc, 0xac, 0x4e, 0xb6, 0x9a, 0x98, 0x32, 0x93, 0x7f, 0x0c, 0x6c,
					0x8d, 0xdd, 0x88, 0x44, 0x42, 0x24, 0x45, 0x73, 0x8a, 0x39, 0x43, 0x64, 0xde, 0x70, 0x4a, 0x46,
				},
			},
			TraceBeginBlock:      1000,
			TraceEndBlock:        2000,
			ChainStateBeginBlock: 8000,
			ChainStateEndBlock:   9000,
		},
	}

	assert.Equal(t, req, expected)
}

var blockResult = ship.Result{
	BlocksResult: &ship.GetBlocksResultV0{
		Head: ship.BlockPosition{
			BlockNum: 302632245,
			BlockID: chain.Checksum256{
				0x12, 0x09, 0xcd, 0x35, 0x19, 0x8a, 0x15, 0x71,
				0x64, 0x8f, 0x66, 0xb2, 0xbe, 0xa9, 0xc0, 0x80,
				0x66, 0x47, 0xd7, 0x88, 0x70, 0xe9, 0xc0, 0x7e,
				0x3c, 0x00, 0x78, 0x87, 0xc7, 0x6a, 0x8b, 0xdf,
			},
		},
		LastIrreversible: ship.BlockPosition{
			BlockNum: 302631909,
			BlockID: chain.Checksum256{
				0x12, 0x09, 0xcb, 0xe5, 0x15, 0x3e, 0xf0, 0xf1,
				0x4a, 0x2d, 0x95, 0xad, 0x80, 0xd0, 0xf5, 0xb4,
				0xa0, 0xf6, 0xaa, 0x48, 0xa4, 0x46, 0xfc, 0x67,
				0xb4, 0xb2, 0x23, 0x66, 0xd8, 0x29, 0xd2, 0x68,
			},
		},
		ThisBlock: &ship.BlockPosition{
			BlockNum: 302632245,
			BlockID: chain.Checksum256{
				0x12, 0x09, 0xcd, 0x35, 0x19, 0x8a, 0x15, 0x71,
				0x64, 0x8f, 0x66, 0xb2, 0xbe, 0xa9, 0xc0, 0x80,
				0x66, 0x47, 0xd7, 0x88, 0x70, 0xe9, 0xc0, 0x7e,
				0x3c, 0x00, 0x78, 0x87, 0xc7, 0x6a, 0x8b, 0xdf,
			},
		},
		PrevBlock: &ship.BlockPosition{
			BlockNum: 302632244,
			BlockID: chain.Checksum256{
				0x12, 0x09, 0xcd, 0x34, 0x53, 0xc5, 0xf6, 0xa6,
				0xbf, 0x57, 0x6e, 0xee, 0x44, 0x3b, 0x4c, 0x94,
				0xb8, 0xff, 0x04, 0x08, 0x50, 0xe9, 0xbb, 0x29,
				0x67, 0x55, 0xa2, 0x56, 0xa8, 0xae, 0x17, 0xdb,
			},
		},
		Block: ship.MustMakeSignedBlockBytes(&ship.SignedBlock{
			SignedBlockHeader: ship.SignedBlockHeader{
				BlockHeader: chain.BlockHeader{
					Timestamp: chain.NewBlockTimestamp(
						time.Date(2024, 4, 10, 22, 49, 53, 500, time.UTC),
					),
					Producer:  chain.N("sentnlagents"),
					Confirmed: 0,
					Previous: chain.Checksum256{
						0x12, 0x09, 0xcd, 0x34, 0x53, 0xc5, 0xf6, 0xa6,
						0xbf, 0x57, 0x6e, 0xee, 0x44, 0x3b, 0x4c, 0x94,
						0xb8, 0xff, 0x04, 0x08, 0x50, 0xe9, 0xbb, 0x29,
						0x67, 0x55, 0xa2, 0x56, 0xa8, 0xae, 0x17, 0xdb,
					},
					TransactionMRoot: chain.Checksum256{
						0xa6, 0xd8, 0xb9, 0x79, 0x3a, 0x49, 0xec, 0x3f,
						0x4d, 0xeb, 0x65, 0x51, 0xfe, 0xd8, 0xcb, 0x70,
						0x1f, 0x41, 0x29, 0x13, 0xc2, 0x4c, 0x6a, 0xf3,
						0x91, 0x99, 0xbe, 0x4f, 0x4a, 0x49, 0xc5, 0x9f,
					},
					ActionMRoot: chain.Checksum256{
						0xef, 0xd7, 0xf1, 0x60, 0x40, 0x8a, 0x81, 0x95,
						0x58, 0x5a, 0x46, 0xba, 0x16, 0x1c, 0x7b, 0xe6,
						0x52, 0x66, 0xf9, 0x27, 0xaf, 0xe3, 0x45, 0x0a,
						0xaa, 0x1f, 0x10, 0x6e, 0x82, 0x50, 0x99, 0xfa,
					},
					ScheduleVersion: 756,
					NewProducersV1: &chain.ProducerSchedule{
						Version: 122,
						Producers: []chain.ProducerKey{
							{
								AccountName:     chain.N("sentnlagents"),
								BlockSigningKey: chain.MustNewPublicKeyFromString("EOS6ejjZgCYwiqaCsJu9aNuefNDA8zYSv7eUR8TkKLus7DHdWTHD8"),
							},
							{
								AccountName:     chain.N("sentnlagents"),
								BlockSigningKey: chain.MustNewPublicKeyFromString("EOS8XX9i6yPTjiFptPNEVcghDoCyZ4gburGWkiSNwVuLWojyyE8Lh"),
							},
						},
					},
					HeaderExtensions: []chain.TransactionExtension{
						{
							Type: 2,
							Data: []byte{0x01, 0x02, 0x03, 0x04},
						},
						{
							Type: 43,
							Data: []byte{0x23, 0x8f, 0x27, 0x83},
						},
					},
				},
				ProducerSignature: chain.MustNewSignatureString("SIG_K1_Kepq3YkvjV4xVe7a1AfSrZK8rzsQ3e4zDtPyVbbjS5sfWQumxjGnzTPoP8kn8BJF8FaVHn4EbhbUq8SStsupzJoRiyoVNs"),
			},
			Transactions: []ship.TransactionReceipt{
				{
					TransactionReceiptHeader: chain.TransactionReceiptHeader{
						Status:               chain.TransactionStatusExecuted,
						CPUUsageMicroSeconds: 887,
						NetUsageWords:        18,
					},
					Trx: ship.Transaction{
						Packed: &chain.PackedTransaction{
							Signatures: []chain.Signature{
								chain.MustNewSignatureString("SIG_K1_KkEZbBin7JFPiq4RUNJ5cvBc6GkUnNS7M348WoGeTKKEJ4gpjzvX5YVVsZ2pikPvxLeQ3VmyVw2kD1scx7bCeyadcBnoVn"),
								chain.MustNewSignatureString("SIG_K1_KVs3Tz8J9EYSbsZTNPnrvQqpP1LPoVYN8gMp4ZFnFJedgwAEgyZktSr5gYCM5HHWPJDTrV2jkGF3go4BhdiqjuCZ6spbNN"),
							},
							Compression:           chain.CompressionNone,
							PackedContextFreeData: []byte{},
							PackedTransaction: []byte{
								0xdf, 0x17, 0x17, 0x66, 0x79, 0xcb, 0x19, 0xd7,
								0x32, 0x60, 0x00, 0x00, 0x00, 0x00, 0x01, 0x90,
								0xe2, 0xa5, 0x1c, 0x5f, 0x25, 0xaf, 0x59, 0x00,
								0x00, 0x00, 0x00, 0x00, 0xe9, 0x4c, 0x44, 0x02,
								0x00, 0x00, 0x90, 0x0c, 0xc0, 0xa9, 0xab, 0x64,
								0x00, 0x00, 0x00, 0x00, 0xe8, 0x8a, 0x14, 0xd6,
								0x30, 0x88, 0x08, 0x21, 0x7c, 0x36, 0x4d, 0xfb,
								0x00, 0x00, 0x00, 0x00, 0xa8, 0xed, 0x32, 0x32,
								0x10, 0x30, 0x88, 0x08, 0x21, 0x7c, 0x36, 0x4d,
								0xfb, 0x84, 0xf9, 0xff, 0x0e, 0x00, 0x01, 0x00,
								0x00, 0x00,
							},
						},
					},
				},
				{
					TransactionReceiptHeader: chain.TransactionReceiptHeader{
						Status:               chain.TransactionStatusExecuted,
						CPUUsageMicroSeconds: 220,
						NetUsageWords:        14,
					},
					Trx: ship.Transaction{
						Packed: &chain.PackedTransaction{
							Signatures:            []chain.Signature{chain.MustNewSignatureString("SIG_K1_K4gAVEXHpnMfUnyZWo5xNDSDj6tEFkWooKCRdp2TYN1fDEB3S9Kzk7BQUsDfVxdscTQZJDMthxxYLsuKxkaps8sa9nrMP4")},
							Compression:           chain.CompressionNone,
							PackedContextFreeData: []byte{},
							PackedTransaction: []byte{
								0xcb, 0x1a, 0x17, 0x66, 0x47, 0xca, 0x3b, 0x3a,
								0x27, 0x8d, 0x00, 0x00, 0x00, 0x00, 0x01, 0x30,
								0xa9, 0xcb, 0xe6, 0xaa, 0xa4, 0x16, 0x90, 0x00,
								0x00, 0x00, 0x20, 0x4d, 0x13, 0xb3, 0xc2, 0x01,
								0x00, 0xa4, 0xe1, 0x00, 0x01, 0x4c, 0x78, 0x71,
								0x00, 0x00, 0x00, 0x00, 0xa8, 0xed, 0x32, 0x32,
								0x10, 0x00, 0xa4, 0xe1, 0x00, 0x01, 0x4c, 0x78,
								0x71, 0x29, 0x58, 0x14, 0x00, 0x00, 0x01, 0x00,
								0x00, 0x00,
							},
						},
					},
				},
			},
			BlockExtensions: []ship.Extension{
				{
					Type: 73,
					Data: []byte{0x3a, 0x8f, 0x73, 0xdd},
				},
			},
		}),
		Traces: ship.MustMakeTransactionTraceArray([]ship.TransactionTrace{
			{
				V0: &ship.TransactionTraceV0{
					ID: chain.Checksum256{
						0x66, 0xa1, 0x90, 0xce, 0x90, 0xe1, 0xa2, 0x9e,
						0x8c, 0xe4, 0x56, 0x86, 0x92, 0x6b, 0x09, 0x5d,
						0x00, 0xe1, 0x75, 0x94, 0x54, 0x80, 0xea, 0x2c,
						0x96, 0xf7, 0xae, 0x6f, 0x5b, 0x52, 0x2b, 0x11,
					},
					Status:        chain.TransactionStatusExecuted,
					CPUUsageUS:    100,
					NetUsageWords: 0,
					Elapsed:       0,
					NetUsage:      0,
					Scheduled:     false,
					ActionTraces: []*ship.ActionTrace{
						{
							V0: &ship.ActionTraceV0{
								ActionOrdinal:        1,
								CreatorActionOrdinal: 0,
								Receipt: &ship.ActionReceipt{
									V0: &ship.ActionReceiptV0{
										Receiver: chain.N("eosio"),
										ActDigest: chain.Checksum256{
											0xf6, 0x4d, 0xf3, 0x66, 0xe8, 0x9e, 0x67, 0xe7,
											0xfd, 0xa7, 0x81, 0x4c, 0x29, 0x6f, 0xf4, 0xd2,
											0x9f, 0xaf, 0x7e, 0x7f, 0x49, 0xad, 0x28, 0x71,
											0x9f, 0xc6, 0x8f, 0xdf, 0x4c, 0xa1, 0x2f, 0x0e,
										},
										GlobalSequence: 89053614934,
										RecvSequence:   471624500,
										AuthSequence: []ship.AccountAuthSequence{
											{
												Account:  chain.N("eosio"),
												Sequence: 370136110,
											},
										},
										CodeSequence: 16,
										ABISequence:  10,
									},
								},
								Receiver: chain.N("eosio"),
								Act: chain.Action{
									Account: chain.N("eosio"),
									Name:    chain.N("onblock"),
									Authorization: []chain.PermissionLevel{
										{
											Actor:      chain.N("eosio"),
											Permission: chain.N("active"),
										},
									},
									Data: []byte{
										0x22, 0xa8, 0x53, 0x5b, 0x80, 0xf3, 0x54, 0xcc,
										0xc4, 0x99, 0xa7, 0xc2, 0x00, 0x00, 0x12, 0x09,
										0xcd, 0x33, 0x23, 0xcf, 0xe6, 0x57, 0x04, 0x7f,
										0xb6, 0xe6, 0x4f, 0xc1, 0x00, 0x7b, 0x82, 0x6f,
										0x16, 0xa0, 0xb3, 0xfe, 0xae, 0x8b, 0x51, 0x17,
										0xe3, 0xbb, 0x5b, 0x4e, 0xd6, 0x6a, 0x3f, 0xd0,
										0xe7, 0x14, 0xb5, 0x90, 0x6b, 0x4c, 0xee, 0x7a,
										0xd8, 0x08, 0xb6, 0x19, 0x68, 0xdd, 0x96, 0xf4,
										0xbc, 0xf7, 0xf5, 0x12, 0x62, 0x95, 0x06, 0xd4,
										0xaa, 0x0f, 0x62, 0x8c, 0x16, 0x56, 0xdb, 0x18,
										0x78, 0x85, 0xd5, 0x5f, 0x96, 0xdf, 0x91, 0xf0,
										0x1c, 0x66, 0x6d, 0xe8, 0xd3, 0x41, 0x59, 0x70,
										0x13, 0x4f, 0x8f, 0xaf, 0x41, 0xb1, 0xeb, 0x3a,
										0x47, 0x55, 0x95, 0x28, 0xfb, 0x83, 0xf4, 0x02,
										0x00, 0x00, 0x00, 0x00,
									},
								},
								ContextFree: true,
								Elapsed:     231,
								Console:     "console",
								AccountRamDeltas: []ship.AccountDelta{
									{
										Account: chain.N("eosio"),
										Delta:   0,
									},
								},
								Except:    "except1",
								ErrorCode: 0xdeadbeef,
							},
						},
					},
					AccountDelta: nil,
					Except:       "except2",
					ErrorCode:    918,
					FailedDtrxTrace: &ship.TransactionTrace{
						V0: &ship.TransactionTraceV0{
							ID: chain.Checksum256{
								0x66, 0xa1, 0x90, 0xce, 0x90, 0xe1, 0xa2, 0x9e,
								0x8c, 0xe4, 0x56, 0x86, 0x92, 0x6b, 0x09, 0x5d,
								0x00, 0xe1, 0x75, 0x94, 0x54, 0x80, 0xea, 0x2c,
								0x96, 0xf7, 0xae, 0x6f, 0x5b, 0x52, 0x2b, 0x11,
							},
							Status:        chain.TransactionStatusExpired,
							CPUUsageUS:    212,
							NetUsageWords: 27,
							Elapsed:       22,
							NetUsage:      88,
							Scheduled:     true,
							ActionTraces: []*ship.ActionTrace{
								{
									V0: &ship.ActionTraceV0{
										ActionOrdinal:        2,
										CreatorActionOrdinal: 4,
										Receipt: &ship.ActionReceipt{
											V0: &ship.ActionReceiptV0{
												Receiver: chain.N("pizzachain11"),
												ActDigest: chain.Checksum256{
													0x52, 0xd8, 0xe2, 0xfe, 0x17, 0x0f, 0xd3, 0x90, 0x53, 0xda, 0x08, 0x42, 0x7b, 0x02, 0x9f, 0x53,
													0x5f, 0xee, 0x1f, 0xda, 0xcd, 0xe2, 0x51, 0x35, 0xc0, 0x03, 0xcb, 0xca, 0x22, 0x92, 0x49, 0x81,
												},
												GlobalSequence: 1728371,
												RecvSequence:   7828,
												AuthSequence: []ship.AccountAuthSequence{
													{
														Account:  chain.N("qubiclesapp1"),
														Sequence: 17267381,
													},
												},
												CodeSequence: 10,
												ABISequence:  24,
											},
										},
										Receiver: chain.N("eosio.token"),
										Act: chain.Action{
											Account: chain.N("eosio.token"),
											Name:    chain.N("transfer"),
											Authorization: []chain.PermissionLevel{
												{
													Actor:      chain.N("qubiclesapp1"),
													Permission: chain.N("active"),
												},
											},
											Data: []byte{0x02, 0x43, 0xde},
										},
										ContextFree: true,
										Elapsed:     2881,
										Console:     "console",
										AccountRamDeltas: []ship.AccountDelta{
											{
												Account: chain.N("qubiclesapp1"),
												Delta:   29,
											},
										},
										Except:    "except3",
										ErrorCode: 6372,
									},
								},
							},
							AccountDelta: &ship.AccountDelta{
								Account: chain.N("eosio"),
								Delta:   -2,
							},
							Except:          "except4",
							ErrorCode:       17821,
							FailedDtrxTrace: nil,
							Partial:         nil,
						},
					},
					Partial: &ship.PartialTransaction{
						V0: &ship.PartialTransactionV0{
							Expiration:       171279471,
							RefBlockNum:      52089,
							RefBlockPrefix:   1613944601,
							MaxNetUsageWords: 0,
							MaxCpuUsageMs:    0,
							DelaySec:         0,
							TransactionExtensions: []ship.Extension{
								{
									Type: 12983,
									Data: []byte{0x0f, 0x02, 0x3f, 0xe3},
								},
							},
							Signatures: []chain.Signature{
								chain.MustNewSignatureString("SIG_K1_KkEZbBin7JFPiq4RUNJ5cvBc6GkUnNS7M348WoGeTKKEJ4gpjzvX5YVVsZ2pikPvxLeQ3VmyVw2kD1scx7bCeyadcBnoVn"),
								chain.MustNewSignatureString("SIG_K1_KVs3Tz8J9EYSbsZTNPnrvQqpP1LPoVYN8gMp4ZFnFJedgwAEgyZktSr5gYCM5HHWPJDTrV2jkGF3go4BhdiqjuCZ6spbNN"),
							},
							ContextFreeData: []byte{0x02, 0x23, 0xfe, 0x00},
						},
					},
				},
			},
			{
				V0: &ship.TransactionTraceV0{
					ID: chain.Checksum256{
						0x69, 0x6d, 0xbb, 0xa1, 0x1c, 0x98, 0x19, 0x9e,
						0x0c, 0x23, 0x0a, 0x62, 0x4f, 0xaa, 0x96, 0x5f,
						0xb8, 0x2e, 0x75, 0x6a, 0x12, 0x87, 0x77, 0x62,
						0xbb, 0x81, 0x87, 0x94, 0xbe, 0xdd, 0x3f, 0x96,
					},
					Status:        chain.TransactionStatusExecuted,
					CPUUsageUS:    220,
					NetUsageWords: 14,
					Elapsed:       0,
					NetUsage:      112,
					Scheduled:     false,
					ActionTraces: []*ship.ActionTrace{
						{
							V1: &ship.ActionTraceV1{
								ActionOrdinal:        1,
								CreatorActionOrdinal: 0,
								Receipt: &ship.ActionReceipt{
									V0: &ship.ActionReceiptV0{
										Receiver: chain.N("m.federation"),
										ActDigest: chain.Checksum256{
											0x36, 0x76, 0xb7, 0x56, 0x9a, 0x7e, 0x35, 0xb4,
											0xc9, 0x05, 0xeb, 0x22, 0x71, 0x17, 0x47, 0x91,
											0x10, 0xcc, 0x4e, 0xf0, 0x70, 0x28, 0x23, 0x8b,
											0xac, 0x0a, 0xaf, 0x65, 0xff, 0xb1, 0x03, 0xf4,
										},
										GlobalSequence: 89053614938,
										RecvSequence:   30152481756,
										AuthSequence: []ship.AccountAuthSequence{
											{
												Account:  chain.N("i5w4s.c.wam"),
												Sequence: 11949,
											},
										},
										CodeSequence: 152,
										ABISequence:  40,
									},
								},
								Receiver: chain.N("m.federation"),
								Act: chain.Action{
									Account: chain.N("m.federation"),
									Name:    chain.N("setland"),
									Authorization: []chain.PermissionLevel{
										{
											Actor:      chain.N("i5w4s.c.wam"),
											Permission: chain.N("active"),
										},
									},
									Data: []byte{
										0x00, 0xa4, 0xe1, 0x00, 0x01, 0x4c, 0x78, 0x71,
										0x29, 0x58, 0x14, 0x00, 0x00, 0x01, 0x00, 0x00,
									},
								},
								ContextFree:      false,
								Elapsed:          231,
								Console:          "console1234",
								AccountRamDeltas: nil,
								Except:           "except5",
								ErrorCode:        0xf3,
								ReturnValue:      []byte{0xbe, 0xef},
							},
						},
					},
					AccountDelta:    nil,
					Except:          "except6",
					ErrorCode:       0xed,
					FailedDtrxTrace: nil,
					Partial:         nil,
				},
			},
		}),
		Deltas: ship.MustMakeTableDeltaArray([]ship.TableDelta{
			{
				V0: &ship.TableDeltaV0{
					Name: "contact_row",
					Rows: []ship.Row{
						{
							Present: true,
							Data: []byte{
								0x00, 0x90, 0xe2, 0xa5, 0x1c, 0x5f, 0x25, 0xaf,
								0x59, 0x90, 0xe2, 0xa5, 0x1c, 0x5f, 0x25, 0xaf,
								0x59, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1c, 0x29,
								0xcd, 0x4b, 0xac, 0x2d, 0x18, 0x00, 0x01, 0x00,
								0x00, 0x90, 0xe2, 0xa5, 0x1c, 0x5f, 0x25, 0xaf,
								0x59, 0x25, 0x4b, 0xac, 0x2d, 0x18, 0x00, 0x01,
								0x00, 0x00, 0xe0, 0x99, 0x09, 0x21, 0x84, 0x10,
								0x50, 0xe7, 0x04, 0x57, 0x6f, 0x6f, 0x64, 0x6e,
								0x1c, 0x03, 0x00, 0x84, 0x03, 0xe1, 0x00, 0xa1,
								0x25, 0x17, 0x66, 0x00, 0x00, 0x00, 0x00,
							},
						},
						{
							Present: true,
							Data: []byte{
								0x00, 0x30, 0xa9, 0xcb, 0xe6, 0xaa, 0xa4, 0x16,
								0x90, 0x30, 0xa9, 0xcb, 0xe6, 0xaa, 0xa4, 0x16,
								0x90, 0x00, 0x00, 0x00, 0x40, 0x61, 0x1d, 0x29,
								0xcd, 0x69, 0xda, 0x68, 0x13, 0x00, 0x01, 0x00,
								0x00, 0x30, 0xa9, 0xcb, 0xe6, 0xaa, 0xa4, 0x16,
								0x90, 0x0c, 0x69, 0xda, 0x68, 0x13, 0x00, 0x01,
								0x00, 0x00, 0x91, 0x17, 0x17, 0x66,
							},
						},
					},
				},
			},
			{
				V0: &ship.TableDeltaV0{
					Name: "contact_row",
					Rows: []ship.Row{
						{
							Present: true,
							Data: []byte{
								0x00, 0x90, 0x15, 0xbc, 0x46, 0x22, 0x27, 0x69,
								0x36, 0x90, 0x15, 0xbc, 0x46, 0x22, 0x27, 0x69,
								0x36, 0x00, 0x00, 0x00, 0x00, 0x00, 0xac, 0xa2,
								0xc1, 0x8e, 0x46, 0x25, 0x09, 0x00, 0x00, 0x00,
								0x00, 0x00, 0x00, 0x90, 0x86, 0x03, 0x6f, 0x69,
								0x8d, 0xcb, 0xd2, 0x93, 0x16, 0xd3, 0xd5, 0x89,
								0x01, 0x68, 0x15, 0x93, 0x60, 0x7f, 0xf5, 0x4f,
								0x6d, 0x08, 0x3a, 0xae, 0x37, 0x05, 0x9c, 0x58,
								0x38, 0x34, 0xd5, 0xde, 0xdc, 0xd9, 0xe7, 0x6f,
								0xc9,
							},
						},
						{
							Present: false,
							Data: []byte{
								0x00, 0x90, 0x15, 0xbc, 0x46, 0x22, 0x27, 0x69,
								0x36, 0x90, 0x15, 0xbc, 0x46, 0x22, 0x27, 0x69,
								0x36, 0x00, 0x00, 0x00, 0x00, 0x00, 0xac, 0xa2,
								0xc1, 0x18, 0x44, 0x25, 0x09, 0x00, 0x00, 0x00,
								0x00, 0x00, 0x00, 0x90, 0x86, 0x03, 0x6f, 0x69,
								0x8d, 0x50, 0x28, 0x62, 0x7b, 0x29, 0x0d, 0xce,
								0x78, 0xa6, 0x68, 0x5a, 0xdc, 0x50, 0xbb, 0x04,
								0x70, 0x13, 0x8f, 0xe7, 0x6d, 0xd6, 0xb7, 0xb8,
								0x8c, 0x28, 0x94, 0x57, 0xe1, 0x50, 0x33, 0x8f,
								0xb6,
							},
						},
					},
				},
			},
		}),
	},
}

func loadHex(filename string) []byte {
	file, err := os.Open(fmt.Sprintf("../testdata/%s.hex", filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := io.ReadAll(hex.NewDecoder(file))
	if err != nil {
		panic(err)
	}
	return data
}

var blockResultEncoded = loadHex("blockresult")

func TestBlocksResultEncode(t *testing.T) {
	buf := new(bytes.Buffer)
	enc := abi.NewEncoder(buf, abi.DefaultEncoderFunc)
	err := enc.Encode(blockResult)
	assert.NoError(t, err)

	data, err := io.ReadAll(buf)
	assert.NoError(t, err)
	assert.Equal(t, data, blockResultEncoded)
}

func TestBlocksResultDecode(t *testing.T) {
	actual := ship.Result{}
	err := abi.NewDecoder(bytes.NewBuffer(blockResultEncoded), abi.DefaultDecoderFunc).Decode(&actual)
	assert.NoError(t, err)

	assert.Equal(t, actual, blockResult)
}
