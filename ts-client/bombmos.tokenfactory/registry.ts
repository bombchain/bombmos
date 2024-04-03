import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryParamsResponse } from "./types/bombmos/tokenfactory/query";
import { QueryGetDenomRequest } from "./types/bombmos/tokenfactory/query";
import { QueryGetDenomResponse } from "./types/bombmos/tokenfactory/query";
import { MsgUpdateParams } from "./types/bombmos/tokenfactory/tx";
import { MsgCreateDenom } from "./types/bombmos/tokenfactory/tx";
import { MsgUpdateDenom } from "./types/bombmos/tokenfactory/tx";
import { MsgUpdateDenomResponse } from "./types/bombmos/tokenfactory/tx";
import { QueryAllDenomRequest } from "./types/bombmos/tokenfactory/query";
import { Denom } from "./types/bombmos/tokenfactory/denom";
import { GenesisState } from "./types/bombmos/tokenfactory/genesis";
import { MsgUpdateParamsResponse } from "./types/bombmos/tokenfactory/tx";
import { MsgCreateDenomResponse } from "./types/bombmos/tokenfactory/tx";
import { QueryParamsRequest } from "./types/bombmos/tokenfactory/query";
import { QueryAllDenomResponse } from "./types/bombmos/tokenfactory/query";
import { Params } from "./types/bombmos/tokenfactory/params";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/bombmos.tokenfactory.QueryParamsResponse", QueryParamsResponse],
    ["/bombmos.tokenfactory.QueryGetDenomRequest", QueryGetDenomRequest],
    ["/bombmos.tokenfactory.QueryGetDenomResponse", QueryGetDenomResponse],
    ["/bombmos.tokenfactory.MsgUpdateParams", MsgUpdateParams],
    ["/bombmos.tokenfactory.MsgCreateDenom", MsgCreateDenom],
    ["/bombmos.tokenfactory.MsgUpdateDenom", MsgUpdateDenom],
    ["/bombmos.tokenfactory.MsgUpdateDenomResponse", MsgUpdateDenomResponse],
    ["/bombmos.tokenfactory.QueryAllDenomRequest", QueryAllDenomRequest],
    ["/bombmos.tokenfactory.Denom", Denom],
    ["/bombmos.tokenfactory.GenesisState", GenesisState],
    ["/bombmos.tokenfactory.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/bombmos.tokenfactory.MsgCreateDenomResponse", MsgCreateDenomResponse],
    ["/bombmos.tokenfactory.QueryParamsRequest", QueryParamsRequest],
    ["/bombmos.tokenfactory.QueryAllDenomResponse", QueryAllDenomResponse],
    ["/bombmos.tokenfactory.Params", Params],
    
];

export { msgTypes }