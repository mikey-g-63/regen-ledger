package cli

//func GetCmdStoreData(cdc *codec.Codec) *cobra.Command {
//	return &cobra.Command{
//		Use:   "store <data>",
//		Short: "store some data on the blockchain",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
//
//			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
//
//			if err := cliCtx.EnsureAccountExists(); err != nil {
//				return err
//			}
//
//			account := cliCtx.GetFromAddress()
//
//			msg := data.NewMsgStoreData([]byte(args[0]), account)
//			err := msg.ValidateBasic()
//			if err != nil {
//				return err
//			}
//
//			cliCtx.PrintResponse = true
//
//			return utils.CompleteAndBroadcastTxCLI(txBldr, cliCtx, []sdk.Msg{msg})
//		},
//	}
//}
