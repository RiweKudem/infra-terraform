// types.ts

import { TerraformBackendConfig } from './terraform-backend-config.type';
import { TerraformVariableConfig } from './terraform-variable-config.type';

export enum TerraformBackendType {
  S3 = 's3',
  GCS = 'gs',
  AzureBlobStorage = 'azurerm',
}

export enum TerraformBackendConfigEncryption {
  HmacSHA256 = 'hmac-sha256',
  HmacSHA256_96 = 'hmac-sha256-96',
  HmacSHA1_96 = 'hmac-sha1-96',
}

export interface TerraformBackendConfig {
  bucket: string;
  key: string;
  region: string;
  access_key: string;
  secret_key: string;
  skip_requesting_account_id: boolean;
  skip_region_validation: boolean;
  assume_role: boolean;
  assume_role_arn: string;
  encrypt: TerraformBackendConfigEncryption;
  encrypt_key: string;
  dynamodb_table: string;
  dynamodb_endpoints: string;
  endpoint: string;
  profile: string;
  shared_credentials_file: string;
  skip_credentialsvalidation: boolean;
  skip_credentialsvalidation: boolean;
  terraform_lock_timeout: number;
  terraform_lock_id: string;
  backend: TerraformBackendType;
}

export interface TerraformVariableConfig {
  type: string;
  default: any;
  description: string;
  sensitive: boolean;
  default: any;
  type: string;
  description: string;
  sensitive: boolean;
  default: any;
  type: string;
  default: any;
  description: string;
  sensitive: boolean;
}

export enum TerraformVariableType {
  String = 'string',
  Number = 'number',
  Boolean = 'boolean',
  List = 'list',
  Tuple = 'tuple',
  Object = 'object',
  Set = 'set',
}

export interface TerraformVariable {
  type: TerraformVariableType;
  default: any;
  description: string;
  sensitive: boolean;
}