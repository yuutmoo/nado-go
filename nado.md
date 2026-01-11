# Table of Contents

- [Overview](#overview)
- [Contracts](#contracts)
  - [📡API](#api)
- [API Changelog](#api-changelog)
- [Archive (indexer)](#archive-indexer)
- [Candlesticks](#candlesticks)
- [Direct Deposit Address](#direct-deposit-address)
- [Edge Candlesticks](#edge-candlesticks)
- [Edge Market Snapshots](#edge-market-snapshots)
- [Events](#events)
- [Fast Withdrawal Signature](#fast-withdrawal-signature)
- [Funding Rate](#funding-rate)
- [Ink Airdrop](#ink-airdrop)
- [Interest & funding payments](#interest-funding-payments)
- [Isolated Subaccounts](#isolated-subaccounts)
- [Linked Signer Rate Limit](#linked-signer-rate-limit)
- [Linked Signers](#linked-signers)
- [Liquidation Feed](#liquidation-feed)
- [Market Snapshots](#market-snapshots)
- [Matches](#matches)
- [NLP Funding Payments](#nlp-funding-payments)
- [NLP Interest Payments](#nlp-interest-payments)
- [NLP Snapshots](#nlp-snapshots)
- [Oracle Price](#oracle-price)
- [Oracle Snapshots](#oracle-snapshots)
- [Orders](#orders)
- [Perp Prices](#perp-prices)
- [Product Snapshots](#product-snapshots)
- [Quote Price](#quote-price)
- [Sequencer Backlog](#sequencer-backlog)
- [Signatures](#signatures)
- [Subaccount Snapshots](#subaccount-snapshots)
- [Subaccounts](#subaccounts)
- [Definitions / Formulas](#definitions-formulas)
- [Depositing](#depositing)
- [🔌Endpoints](#endpoints)
- [Errors](#errors)
- [Gateway](#gateway)
- [Executes](#executes)
- [Burn NLP](#burn-nlp)
- [Cancel And Place](#cancel-and-place)
- [Cancel Orders](#cancel-orders)
- [Cancel Product Orders](#cancel-product-orders)
- [Link Signer](#link-signer)
- [Liquidate Subaccount](#liquidate-subaccount)
- [Mint NLP](#mint-nlp)
- [Place Order](#place-order)
- [Place Orders](#place-orders)
- [Transfer Quote](#transfer-quote)
- [Withdraw Collateral](#withdraw-collateral)
- [Queries](#queries)
- [All Products](#all-products)
- [Edge All Products](#edge-all-products)
- [Fee Rates](#fee-rates)
- [Health Groups](#health-groups)
- [Insurance](#insurance)
- [Isolated Positions](#isolated-positions)
- [Linked Signer](#linked-signer)
- [Market Liquidity](#market-liquidity)
- [Market Prices](#market-prices)
- [Max NLP Burnable](#max-nlp-burnable)
- [Max NLP Mintable](#max-nlp-mintable)
- [Max Order Size](#max-order-size)
- [Max Withdrawable](#max-withdrawable)
- [NLP Locked Balances](#nlp-locked-balances)
- [NLP Pool Info](#nlp-pool-info)
- [Nonces](#nonces)
- [Order](#order)
- [Status](#status)
- [Subaccount Info](#subaccount-info)
- [Symbols](#symbols)
- [Signing](#signing)
- [Examples](#examples)
- [Q&A](#q-a)
- [Integrate via Smart Contracts](#integrate-via-smart-contracts)
- [Order Appendix](#order-appendix)
- [Rate limits](#rate-limits)
- [Subscriptions](#subscriptions)
- [Authentication](#authentication)
- [Streams](#streams)
- [Trigger](#trigger)
- [List Trigger Orders](#list-trigger-orders)
- [List TWAP Executions](#list-twap-executions)
- [V2](#v2)
- [APR](#apr)
- [Assets](#assets)
- [Orderbook](#orderbook)
- [Pairs](#pairs)
- [Tickers](#tickers)
- [Trades](#trades)
- [Withdrawing (on-chain)](#withdrawing-on-chain)
- [🚀Get Started](#get-started)
- [🔧Common Errors](#common-errors)
- [📖Core Concepts](#core-concepts)
- [💰First Deposit](#first-deposit)
- [🔗Linked Signers](#linked-signers)
- [⚡Quickstart](#quickstart)
  - [📘TypeScript SDK](#typescript-sdk)
- [Getting Started](#getting-started)
- [How To](#how-to)
- [Create a Nado client](#create-a-nado-client)
- [Deposit Funds](#deposit-funds)
- [Manage Orders](#manage-orders)
- [Query Markets & Products](#query-markets-products)
- [Useful Common Functions](#useful-common-functions)
- [Withdraw Funds](#withdraw-funds)
- [FAQs](#faqs)
- [Fees & Rebates](#fees-rebates)
- [Funding Rates](#funding-rates)
- [Legal](#legal)
- [Restricted Territories](#restricted-territories)
- [Liquidations](#liquidations)
- [Maintenance Windows](#maintenance-windows)
- [Margin Types](#margin-types)
- [Market Parameters](#market-parameters)
- [Mission](#mission)
- [NLP](#nlp)
- [Onboarding Tutorial](#onboarding-tutorial)
- [Bridging USDT0 to Ink](#bridging-usdt0-to-ink)
- [Oracles](#oracles)
- [Order Types](#order-types)
- [Orderbook Architecture](#orderbook-architecture)
- [PnL Settlements](#pnl-settlements)
- [Products](#products)
- [Subaccounts & Health](#subaccounts-health)

---


# Overview

Source: https://docs.nado.xyz/

Nado is the central-limit orderbook DEX where raw power collides with surgical precision to unleash breathtaking performance for traders. Trade spot and perpetuals via unified margin, where kinetic collateral dynamically flows across positions — netting risks in real-time for superior capital efficiency.

Built on Ink and backed by the team that brought you Kraken, Nado delivers a feature-rich experience that cuts through noise to amplify your trading edge.

Welcome to the perfect storm.

---

## Key Features of Nado

## Trade Spot & Perpetual Markets

> Seamlessly buy, sell, and leverage assets with deep liquidity and blazing-fast speeds. Nado supports majors like BTC and ETH, enabling up to 20x leverage on perpetuals.

> Multiple collateral options and margin trading unlock flexibility, ensuring a precise, comprehensive experience.

## Nado Engine: Speed & Precision Redefined

> Nado’s raw power converts split-second opportunities into an edge over the competition.

> At Nado's core is the high-performance off-chain sequencer and on-chain risk engine with settlement via the Ink L2. Incoming trades strike the orderbook in a flash at 5-15 ms latency and settle on-chain in batches. Plugging into the HFT-friendly API delivers pure, programmable power when it matters most.

> Nado introduces an unmatched on-chain trading experience that empowers traders' instincts to execute their strategies with exacting precision and control.

## Unified Margin

> Harness your full portfolio as collateral across spot, perpetuals, and money markets in one unified margin account. Collateral in constant motion puts your portfolio to work with dynamic margin offsets that reduce overall margin requirements, maximize capital efficiency, and amplify buying power.

> For example, a spot ETH hold offsets a BTC perps short, liabilities consolidate intuitively, and traders can maximize their capital efficiency. Conduct delta neutral strategies without worrying about liquidation and scale your basis trades with higher capital efficiency.

> Trade smarter with leveraged spot, multiple order types, and all funds – deposits, positions, PnL – working toward your margin.

## Money Markets

> Earn yield automatically on deposits, borrow against margin with flexible collaterals like wETH or USDT0, and turbo-loop positions effortlessly with spot margin trading.

> The margin engine automatically recognizes spread trades as hedges so you can scale your basis trades with ease while your margin levels automatically adjust in the background.

## NLP: Passive Yields from Altcoin Edges

> Deposit USDT0 into Nado's Liquidity Provider (NLP) vault and unleash it as a passive force – a sleek counterparty that deepens perpetual liquidity while capturing yields from the chaos of crypto markets.

> Powered by sub-vault strategies, such as active liquidations that buy positions at discounts and exit swiftly for profit, NLP funnels capital into perpetuals-only plays.

> Traders gain tighter spreads and faster fills while LPs earn passively – bolstering liquidity for everyone on the platform.

## Super Low Fees

> Nado's low-fee model strips away exchange drag, with taker fees as low as 1.5 bps and maker rebates scaling up to -0.8 bps across volume-based tiers. Trade more, pay less.

## Pro Tools for Pro Traders

> Advanced tools meet trader instinct. One-click basis trades, TWAPs, scale orders, and conditional triggers are all available at your fingertips in one streamlined interface. Deploy your favorite trading strategies with tools built for pros and newcomers alike.

## Self-Custody

> Your keys, your storm. Retain full control over assets amidst the chaos, trading with DEX transparency and CEX precision – mitigating centralized risks.

## One DEX. Every Market.

> Cryptos, equities, commodities, private equities, and more coming soon.

## Nado: The Perfect Storm

> Amidst the chaos of crypto, precision reigns. Don't just endure volatility’s chaos, make it yours to command on Nado.

---

[NextMission](/mission)

Last updated 17 days ago

---


# Contracts

Source: https://docs.nado.xyz/contracts

> **Nado Github Repo**: <https://github.com/nadohq/nado-contracts>

## Mainnet Contracts — Ink

> **Explorer URL**: [https://explorer.inkonchain.com](https://explorer.inkonchain.com/)

Contract Name

Address

**Deployer**

[0xC1cC56caB60e832665E6c3780BfEBe3C1C971603](https://explorer.inkonchain.com/address/0xC1cC56caB60e832665E6c3780BfEBe3C1C971603)

**Quote**

[0x0200C29006150606B650577BBE7B6248F58470c1](https://explorer.inkonchain.com/address/0x0200C29006150606B650577BBE7B6248F58470c1)

**Querier**

[0x68798229F88251b31D534733D6C4098318c9dff8](https://explorer.inkonchain.com/address/0x68798229F88251b31D534733D6C4098318c9dff8)

**Clearinghouse**

[0xD218103918C19D0A10cf35300E4CfAfbD444c5fE](https://explorer.inkonchain.com/address/0xD218103918C19D0A10cf35300E4CfAfbD444c5fE)

**Endpoint**

[0x05ec92D78ED421f3D3Ada77FFdE167106565974E](https://explorer.inkonchain.com/address/0x05ec92D78ED421f3D3Ada77FFdE167106565974E)

**SpotEngine**

[0xFcD94770B95fd9Cc67143132BB172EB17A0907fE](https://explorer.inkonchain.com/address/0xFcD94770B95fd9Cc67143132BB172EB17A0907fE)

**PerpEngine**

[0xF8599D58d1137fC56EcDd9C16ee139C8BDf96da1](https://explorer.inkonchain.com/address/0xF8599D58d1137fC56EcDd9C16ee139C8BDf96da1)

**WithdrawPool**

[0x09fb495AA7859635f755E827d64c4C9A2e5b9651](https://explorer.inkonchain.com/address/0x09fb495AA7859635f755E827d64c4C9A2e5b9651)

---

## Testnet Contracts — Ink Sepolia

> **Explorer URL**: <https://explorer-sepolia.inkonchain.com/>

**Deployer**

[0x59841b3761Ed1D089a783E4d7CB49E4534CD4F85](https://explorer-sepolia.inkonchain.com/address/0x59841b3761Ed1D089a783E4d7CB49E4534CD4F85)

**Quote**

[0x60F50F902b2E91aef7D6c700Eb22599e297fa86F](https://explorer-sepolia.inkonchain.com/address/0x60F50F902b2E91aef7D6c700Eb22599e297fa86F)

**Querier**

[0x8E693BEa316bcC0F4f8be403081b954a0E3743C8](https://explorer-sepolia.inkonchain.com/address/0x8E693BEa316bcC0F4f8be403081b954a0E3743C8)

**Clearinghouse**

[0x23a283B359D55A941bBeEC58801B6b17D955CC73](https://explorer-sepolia.inkonchain.com/address/0x23a283B359D55A941bBeEC58801B6b17D955CC73)

**Endpoint**

[0x698D87105274292B5673367DEC81874Ce3633Ac2](https://explorer-sepolia.inkonchain.com/address/0x698D87105274292B5673367DEC81874Ce3633Ac2)

**SpotEngine**

[0x3352b2fF0fAc4ce38A6eA1C188cF4F924df54E5D](https://explorer-sepolia.inkonchain.com/address/0x3352b2fF0fAc4ce38A6eA1C188cF4F924df54E5D)

**PerpEngine**

[0x4E859C47fea3666B5053B16C81AF64e77567702e](https://explorer-sepolia.inkonchain.com/address/0x4E859C47fea3666B5053B16C81AF64e77567702e)

**WithdrawPool**

[0xBD672Fe513acbA5c1ceE7b02F998A1B542852b3b](https://explorer-sepolia.inkonchain.com/address/0xBD672Fe513acbA5c1ceE7b02F998A1B542852b3b)


Last updated 1 month ago

---


# 📡API

Source: https://docs.nado.xyz/developer-resources/api

## Overview

Nado's API is divided into the following categories:

1. A **websocket/REST** API (**gateway**) that supports writes (executes) and polling (queries).
2. A **subscriptions** API that allows to subscribe to live data feeds.
3. An **indexer** API (**archive**) that allows you to query historical data.
4. A **trigger** API that allows to execute orders only under specified price conditions.

[Gateway](/developer-resources/api/gateway)[Subscriptions](/developer-resources/api/subscriptions)

Last updated 17 days ago

---


# API Changelog

Source: https://docs.nado.xyz/developer-resources/api/api-changelog

This document tracks all changes to the Nado API.

---

**December 22, 2025**

**Archive Indexer: Orders & Matches fields**

* Added `closed_amount` (x18) and `realized_pnl` (x18) to archive indexer responses for both [orders](/developer-resources/api/archive-indexer/orders) and [matches](/developer-resources/api/archive-indexer/matches).

---

**December 11, 2025**

**Risk System Updates**

**Spread Weight Caps**

* Introduced upper bounds for spread weights to manage risk at extreme leverage levels:

  + `initial_spread_weight`: Maximum **0.99**
  + `maintenance_spread_weight`: Maximum **0.994**
* **Impact**:

  + Existing markets (≤20x leverage): No change in behavior
  + Future high-leverage markets (30x+): Spread positions will have capped health benefits
  + Prevents extreme leverage abuse via spread positions
* **Technical Details**:

  + Base spread weight calculated as: `spread_weight = 1 - (1 - product_weight) / 5`
  + Final spread weight: `min(spread_weight, cap)`
  + Cap applies during health calculations for spread positions

**Minimum Liquidation Penalties**

* Introduced minimum distance requirements between oracle price and liquidation price:

  + **Non-spread liquidations**: Minimum **0.5%** from oracle price
  + **Spread liquidations**: Minimum **0.25%** from oracle price
* **Impact**:

  + Ensures liquidators always have sufficient incentive to execute liquidations
  + Prevents unprofitable liquidation scenarios for low-volatility assets
  + Particularly important for high-leverage positions where natural penalties may be very small
* **Technical Details**:

  + **Non-spread longs**: `oracle_price × (1 - max((1 - maint_asset_weight) / 5, 0.005))`
  + **Non-spread shorts**: `oracle_price × (1 + max((maint_liability_weight - 1) / 5, 0.005))`
  + **Spread selling**: `spot_price × (1 - max((1 - perp_maint_asset_weight) / 10, 0.0025))`
  + **Spread buying**: `spot_price × (1 + max((spot_maint_liability_weight - 1) / 10, 0.0025))`

**API Response Changes**

* No breaking changes to API response structure
* Health calculations and liquidation prices automatically reflect new risk parameters

**Documentation Updates**

* See [Subaccounts & Health](/subaccounts-and-health#spreads) for spread weight cap details
* See [Liquidations](/liquidations#liquidation-price) for minimum liquidation penalty details

---

**December 1, 2025**

**Query Enhancements**

**Pre-State Simulation for SubaccountInfo Query**

* Added `pre_state` parameter to `SubaccountInfo` query

  + Type: `string` (accepts `"true"` or `"false"`)
  + When set to `"true"` along with `txns`, returns a `pre_state` object in the response
  + `pre_state` contains the subaccount state **before** the simulated transactions were applied
  + Useful for comparing before/after states when simulating trades
  + `pre_state` includes:

    - `healths`: Health information before transactions
    - `health_contributions`: Per-product health contributions before transactions
    - `spot_balances`: Spot balances before transactions
    - `perp_balances`: Perpetual balances before transactions

**Use Cases:**

* Position simulation and preview
* Risk analysis for potential trades
* UI/UX for showing before/after comparisons
* Testing transaction impacts without on-chain execution

**Documentation:** See [Subaccount Info Query](/developer-resources/api/gateway/queries/subaccount-info#example-with-pre_state) for detailed examples.

---

**November 20, 2025 - Initial Launch**

## Core Changes

**1. Removal of LP Functionality**

* `SubaccountInfo` no longer has:

  + `lp_balance` in `spot_balances` and `perp_balances`
  + `lp_state` in `spot_products` and `perp_products`
  + `lp_spread_x18` in `book_info` of both `spot_products` and `perp_products`
* Historical `events` no longer include:

  + `net_entry_lp_unrealized`
  + `net_entry_lp_cumulative`

**2. Removal of Redundant Fields**

* `SubaccountInfo` no longer has:

  + `last_cumulative_multiplier_x18` in `balance` of `spot_balances`

**3. Products Config Model Updates**

* Added: `withdraw_fee_x18` and `min_deposit_rate_x18` to `spot_products.config`

**4. Products Risk Model Updates**

* Added: `price_x18` to both `spot_products.risk` and `perp_products.risk`
* Removed: `large_position_penalty_x18`

**5. Deposit Rate Query**

* Removed: `min_deposit_rates` query
* Use `min_deposit_rate_x18` in `spot_products.config` instead

## Market Structure Changes

**6. Removal of Virtual Books**

* `Contracts` query no longer returns `book_addrs`
* `PlaceOrder` verify contract is now `address(product_id)`
  &#xNAN;*Example: product* *18* *→* `0x0000000000000000000000000000000000000012`

**7. Minimum Size denomination**

* `min_size` is now **USDT0 denominated** (not base denominated)

  + `min_size = 10` → minimum order size = 10 USDT0 (`order_price * order_amount`)
* `size_increment` remains **base denominated**

  + Example: BTC with `size_increment = 0.0001` and `min_size = 20`:

    - ✅ Valid: 100,000 \* 0.0002 = 20 USDT0
    - ❌ Invalid: 100,000 \* 0.0001 = 10 USDT0
    - ❌ Invalid: 100,000 \* 0.00025 (not multiple of 0.0001)

## Orders & Signing

**8. Place Orders Execute**

* Added: `place_orders` execute - place multiple orders in a single request

  + Accepts array of orders with same structure as `place_order`
  + Optional `stop_on_failure` parameter to stop processing remaining orders on first failure
  + Returns array of results with `digest` (if successful) or `error` (if failed) for each order
  + Rate limit weight calculated per order

See [Place Orders](/developer-resources/api/gateway/executes/place-orders) for details.

**9. EIP712** `Order` **Struct Update**

Copy

```
struct Order {
    bytes32 sender;
    int128 priceX18;
    int128 amount;
    uint64 expiration;
    uint64 nonce;
    uint128 appendix;
}
```

* New field: `appendix`
* All order flags (IOC, post only, reduce-only, triggers) moved into `appendix`
* `expiration` is now strictly a timestamp
* `appendix` bitfield:

Copy

```
| value   | reserved | trigger | reduce only | order type | isolated | version |
| 64 bits | 50 bits  | 2 bits  | 1 bit       | 2 bits     | 1 bit    | 8 bits  |
```

* Special encodings:

  + `trigger` = 2 or 3 → `value` encodes TWAP settings (`times`, `slippage_x6`)
  + `isolated = 1` → `value` encodes isolated margin
* Constraints:

  + Isolated orders cannot be TWAP
  + TWAP orders must use IOC execution type

See [Order Appendix Docs](/developer-resources/api/order-appendix).

**10. TWAP Order Execution**

* Added `list_twap_executions` query to trigger service
* TWAP orders track individual execution status (pending, executed, failed, cancelled)
* TWAP execution statuses include execution time and engine response data

**11. Trigger Service Rate Limits**

* Updated trigger order limits from 100 pending orders per subaccount to `25 pending orders per product per subaccount`

**12. EIP712 Domain Change**

* Signing domain updated from `Vertex` **→** `Nado`
  See [Signing Docs](/developer-resources/api/gateway/signing).

## Query Updates

**13.** `max_order_size`

* Added: `isolated` parameter - when set to `true`, calculates max order size for an isolated margin position. Defaults to `false`.

**14.** `orders` **Query**

* Added: `trigger_types` parameter - filter orders by trigger type(s)

**15. Historical Events**

* Added: `quote_volume_cumulative` - tracks cumulative trading volume for the subaccount in quote units

  + Available in: `events` and `subaccount_snapshots` queries

**16.** `subaccount_snapshots` **Query**

* Added: `active` parameter - filter snapshots by position status

  + `true`: returns only products with **non-zero balance** at the timestamp
  + `false`: returns products with **event history** before the timestamp (default)

**17. Trigger Orders**

* Added: `place_at` field - timestamp when trigger order should be placed

**18. Removal of** `summary` **Query**

* Removed: `summary` query from indexer API
* Use `subaccount_snapshots` query instead for historical subaccount data

**19. Query Renaming**

* Renamed: `usdc_price` → `quote_price` query

  + See [Quote Price](/developer-resources/api/archive-indexer/quote-price)

**20. Multi-Subaccount** `events`**,** `matches`**,** `orders`

* The indexer `events`, `matches`, and `orders` queries now accept a `subaccounts` array so you can fetch history for multiple subaccounts in a single request instead of fanning out per subaccount. Please note that the old single-subaccount version is **no longer supported**.

## Streams

See [Subscriptions > Streams](/developer-resources/api/subscriptions/streams) for more details

**21.** `OrderUpdate`

* Can now subscribe across all products by setting `product_id = null`
* `product_id` type changed from `u32` → `Option<u32>`

**22.** `Fill`

* Added: `fee`, `submission_idx`, and `appendix`
* Can now subscribe across all products by setting `product_id = null`

**23.** `PositionChange`

* Can now subscribe across all products by setting `product_id = null`
* `product_id` type changed from `u32` → `Option<u32>`
* Added: `isolated` - indicates whether the position change is for an isolated margin position

**24.** `FundingPayment`

* New stream: `FundingPayment`
* Param: `product_id: u32`
* Emits hourly funding payment events

**Request**

Copy

```
{
  "method": "subscribe",
  "stream": {
    "type": "funding_payment",
    "product_id": 1
  },
  "id": 123
}
```

**Response**

Copy

```
{
  "type": "funding_payment",
  "timestamp": 1234567890000,
  "product_id": 1,
  "payment_amount": "1000000000000000000",
  "open_interest": "50000000000000000000",
  "cumulative_funding_long_x18": "100000000000000000",
  "cumulative_funding_short_x18": "-100000000000000000",
  "dt": 3600000
}
```

**25.** `Liquidation`

* New stream: `Liquidation`
* Param: `product_id` or `null` (all products)
* Emits liquidation info (liquidator, liquidatee, amount, price)

**Request**

Copy

```
{
  "method": "subscribe",
  "stream": {
    "type": "liquidation",
    "product_id": 1
  },
  "id": 123
}
```

**Response**

Copy

```
{
  "type": "liquidation",
  "timestamp": "1234567890000",
  "product_ids": [1],
  "liquidator": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
  "liquidatee": "0x8b6fd3859f7065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
  "amount": "1000000000000000000",
  "price": "50000000000000000000"
}
```

**26.** `LatestCandlestick`

* New stream: `LatestCandlestick`
* Params: `product_id`, `granularity` (seconds)
* Emits candlestick updates on every trade

**Request**

Copy

```
{
  "method": "subscribe",
  "stream": {
    "type": "latest_candlestick",
    "product_id": 1,
    "granularity": 60
  },
  "id": 123
}
```

**Response**

Copy

```
{
  "type": "latest_candlestick",
  "timestamp": 1234567890000,
  "product_id": 1,
  "granularity": 60,
  "open_x18": "50000000000000000000",
  "high_x18": "51000000000000000000",
  "low_x18": "49000000000000000000",
  "close_x18": "50500000000000000000",
  "volume": "1000000000000000000"
}
```

**27.** `FundingRate`

* New stream: `FundingRate`
* Param: `product_id` or `null` (all products)
* Emits funding rate updates every 20 seconds
* `funding_rate_x18` and `update_time` values are identical to those from the [Funding Rate](/developer-resources/api/archive-indexer/funding-rate) indexer endpoint

**Request**

Copy

```
{
  "method": "subscribe",
  "stream": {
    "type": "funding_rate",
    "product_id": 1
  },
  "id": 123
}
```

**Subscribe to all products:**

Copy

```
{
  "method": "subscribe",
  "stream": {
    "type": "funding_rate",
    "product_id": null
  },
  "id": 123
}
```

**Response**

Copy

```
{
  "type": "funding_rate",
  "timestamp": "1234567890123456789",
  "product_id": 1,
  "funding_rate_x18": "50000000000000000",
  "update_time": "1234567890"
}
```


Last updated 16 days ago

---


# Archive (indexer)

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer

Using Nado's indexer API you can access historical data in the platform as it is processed by our offchain sequencer. This includes: trading activity, events, candlesticks and more.

You can interact with our indexer by sending `HTTP` requests at `POST [ARCHIVE_ENDPOINT]` alongside a json payload of the query. Endpoints:

`HTTP` requests must set the `Accept-Encoding` to include `gzip`, `br` or `deflate`

## Endpoints

## Testnet:

* `https://archive.test.nado.xyz/v1`

## Available Queries:

[Orders](/developer-resources/api/archive-indexer/orders)[Matches](/developer-resources/api/archive-indexer/matches)[Events](/developer-resources/api/archive-indexer/events)[Candlesticks](/developer-resources/api/archive-indexer/candlesticks)[Edge Candlesticks](/developer-resources/api/archive-indexer/edge-candlesticks)[Product Snapshots](/developer-resources/api/archive-indexer/product-snapshots)[Funding Rate](/developer-resources/api/archive-indexer/funding-rate)[Interest & funding payments](/developer-resources/api/archive-indexer/interest-and-funding-payments)[Oracle Price](/developer-resources/api/archive-indexer/oracle-price)[Oracle Snapshots](/developer-resources/api/archive-indexer/oracle-snapshots)[Perp Prices](/developer-resources/api/archive-indexer/perp-prices)[Market Snapshots](/developer-resources/api/archive-indexer/market-snapshots)[Edge Market Snapshots](/developer-resources/api/archive-indexer/edge-market-snapshots)[Subaccounts](/developer-resources/api/archive-indexer/subaccounts)[Subaccount Snapshots](/developer-resources/api/archive-indexer/subaccount-snapshots)[Linked Signers](/developer-resources/api/archive-indexer/linked-signers)[Linked Signer Rate Limit](/developer-resources/api/archive-indexer/linked-signer-rate-limit)[Isolated Subaccounts](/developer-resources/api/archive-indexer/isolated-subaccounts)[Signatures](/developer-resources/api/archive-indexer/signatures)[Fast Withdrawal Signature](/developer-resources/api/archive-indexer/fast-withdrawal-signature)[NLP Funding Payments](/developer-resources/api/archive-indexer/nlp-funding-payments)[NLP Interest Payments](/developer-resources/api/archive-indexer/nlp-interest-payments)[NLP Snapshots](/developer-resources/api/archive-indexer/nlp-snapshots)[Liquidation Feed](/developer-resources/api/archive-indexer/liquidation-feed)[Sequencer Backlog](/developer-resources/api/archive-indexer/sequencer-backlog)[Direct Deposit Address](/developer-resources/api/archive-indexer/direct-deposit-address)

Last updated 20 days ago

---


# Candlesticks

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/candlesticks

## Rate limits

* Dynamic based on `limit` param provided (**weight = 1 + limit / 20**)

  + E.g: With `limit=100`, you can make up to 400 requests per min or 66 requests / 10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Available Granularities

The following granularities / periods are supported (in seconds):

Granularity name

Granularity value (in seconds)

1 minute

60

5 minutes

300

15 minutes

900

1 hour

3600

2 hours

7200

4 hours

14400

1 day

86400

1 week

604800

4 weeks

2419200

## Request

Product candlesticks

Query product candlesticks ordered by `timestamp` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of product to fetch candlesticks for.

granularity

number

Yes

Granularity value in seconds.

max\_time

number / string

No

When providing `max_time` (unix epoch in seconds), only return candlesticks with timestamp <= `max_time`

limit

number

No

Max number of candlesticks to return. defaults to `100`. max possible of `500`.

## Response

## Response Fields

Field name

Description

submission\_idx

Id of the latest recorded transaction that contributes to the candle.

product\_id

Id of product candle is associated to.

granularity

Candle time interval, expressed in seconds, representing the aggregation period for trading volume and price data

open\_x18

The first fill price of the candle, multiplied by 10^18

high\_x18

The highest recorded fill price during the defined interval of the candle, multiplied by 10^18

low\_x18

The lowest recorded fill price during the defined interval of the candle, multiplied by 10^18

close\_x18

The last price of the candle, multiplied by 10^18

volume

Asset volume, which represents the absolute cumulative fill amounts during the time interval of the candle, multiplied by 10^18


Last updated 1 month ago

---


# Direct Deposit Address

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/direct-deposit-address

## Rate limits

* 240 requests/min or 40 requests/10secs per IP address. (**weight = 10**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Direct Deposit Address

Query the unique direct deposit address for a subaccount.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "direct_deposit_address": {
    "subaccount": "0x79cc76364b5fb263a25bd52930e3d9788fcfeea864656661756c740000000000"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccount

string

Yes

Hex string of the subaccount to fetch the direct deposit address for.

## Response

## Response Fields

## Direct Deposit Address

Field name

Description

subaccount

Hex string of the subaccount

deposit\_address

Unique deposit address for this subaccount

created\_at

Unix epoch time in seconds when the deposit address was created

Direct deposit addresses allow users to deposit funds directly to their subaccount without needing to interact with the smart contract. Funds sent to this address will automatically be credited to the associated subaccount.


Last updated 1 month ago

---


# Edge Candlesticks

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/edge-candlesticks

## Rate limits

* Dynamic based on `limit` param provided (**weight = 1 + limit / 20**)

  + E.g: With `limit=100`, you can make up to 400 requests per min or 66 requests / 10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Available Granularities

The following granularities / periods are supported (in seconds):

Granularity name

Granularity value (in seconds)

1 minute

60

5 minutes

300

15 minutes

900

1 hour

3600

2 hours

7200

4 hours

14400

1 day

86400

1 week

604800

4 weeks

2419200

## Request

Ede candlesticks

Query edge candlesticks ordered by `timestamp` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of product to fetch candlesticks for.

granularity

number

Yes

Granularity value in seconds.

max\_time

number / string

No

When providing `max_time` (unix epoch in seconds), only return candlesticks with timestamp <= `max_time`

limit

number

No

Max number of candlesticks to return. defaults to `100`. max possible of `500`.

## Response

## Response Fields

Field name

Description

submission\_idx

Id of the latest recorded transaction that contributes to the candle.

product\_id

Id of product candle is associated to.

granularity

Candle time interval, expressed in seconds, representing the aggregation period for trading volume and price data

open\_x18

The first fill price of the candle, multiplied by 10^18

high\_x18

The highest recorded fill price during the defined interval of the candle, multiplied by 10^18

low\_x18

The lowest recorded fill price during the defined interval of the candle, multiplied by 10^18

close\_x18

The last price of the candle, multiplied by 10^18

volume

Asset volume, which represents the absolute cumulative fill amounts during the time interval of the candle, multiplied by 10^18


Last updated 1 month ago

---


# Edge Market Snapshots

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/edge-market-snapshots

## Rate limits

**Dynamic based on interval.count.**

* IP weight = `(interval.count.min(500) / 20) + (interval.count.clamp(2, 20) * 2)`
* Scales mainly with interval count.

  + Example: `interval.count=500 → weight=65`, `interval.count=100 → weight=45`.
* Minimum weight per request is `4`.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Market snapshots

Query market snapshots ordered by `timestamp` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
    "edge_market_snapshots": {
        "interval": {
          "count": 2,
          "granularity": 3600,
          "max_time": 1691083697,
        },
        "product_ids": [1, 2]
    }
}
```

## Request Parameters

Parameter

Type

Required

Description

interval

object

Yes

Object to specify desired time period for data

interval.count

number

Yes

Number of snapshots to return, limit 100. Also limited to `interval.count * # product_ids < 2000`

interval.granularity

number

Yes

Granularity value in seconds

interval.max\_time

number / string

No

When providing `max_time` (unix epoch in seconds), only return snapshots with timestamp <= `max_time`. If no value is entered, `max_time` defaults to the current time.

product\_ids

number[]

No

list of product ids to fetch snapshots for, defaults to all products

## Response

**Note**:

* Returns a mapping of `chain_id -> snapshots`

## Response Fields

## Snapshots

**Note**: For product specific fields (i.e. cumulative\_volume, open\_interests), the value is an object which maps product\_ids to their corresponding values.

Field name

Description

timestamp

Timestamp of the snapshot. This may not be perfectly rounded to the granularity since it uses the nearest transaction timestamp less than or equal to `max_time`

cumulative\_users

The cumulative number of subaccounts on Nado. It is updated daily at 9AM ET for historical counts. For current day counts, it is updated every hour.

daily\_active\_users

Daily active users count, updated daily at 9AM ET for historical counts. For current day counts, it is updated every hour.

cumulative\_trades

A map of product\_id -> the cumulative number of trades for the given product\_id.

cumulative\_volumes

A map of product\_id -> cumulative volumes in USDT0 units.

cumulative\_trade\_sizes

A map of product\_id -> cumulative trade sizes in base token

cumulative\_taker\_fees

A map of product\_id -> cumulative taker fees. Taker fees include sequencer fees.

cumulative\_sequencer\_fees

A map of product\_id -> cumulative sequencer fees.

cumulative\_maker\_fees

A map of product\_id -> cumulative maker rebates.

cumulative\_liquidation\_amounts

A map of product\_id -> cumulative liquidation amounts in USDT0 units.

open\_interests

A map of product\_id -> open interests in USDT0 units.

total\_deposits

A map of product\_id -> total deposits held by Nado for a given product at the given time in the base token units.

total\_borrows

A map of product\_id -> total borrows lent by Nado for a given product at the given time in the base token units.

funding\_rates

A map of product\_id -> **hourly** historical funding rates, value returned as **decimal rates** (% = rate \* 100), derived from funding payment amounts. Requires a minimum granularity of 3600 to see non-zero funding rates. Use a granularity where granularity % 3600 = 0 for best results.

deposit\_rates

A map of product\_id -> **daily** deposit rates, values returned as **decimal rates** (% = rate \* 100).

borrow\_rates

A map of product\_id -> **daily** borrow rates, values returned as **decimal rates** (% = rate \* 100).

cumulative\_inflows

A map of product\_id -> cumulative inflows a.k.a deposits in base token units.

cumulative\_outflows

A map of product\_id -> cumulative outflows a.k.a withdraws in base token units.

tvl

The total value locked in USD.


Last updated 1 month ago

---


# Events

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/events

## Rate limits

* IP weight = `2 + (limit * subaccounts.length / 10)` where `limit` defaults to 100 (max 500) and `subaccounts.length` defaults to 1

  + E.g: With `limit=100` and 1 subaccount, weight = 12, allowing up to 200 requests per min or 33 requests / 10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Available Events

Each event corresponds to a transaction type in Nado. See below available events and their `event_type` mapping:

Event Name

Event Type Value

`LiquidateSubaccount`

liquidate\_subaccount

`DepositCollateral`

deposit\_collateral

`WithdrawCollateral`

withdraw\_collateral

`SettlePnl`

settle\_pnl

`MatchOrders`

match\_orders

`MintLp`

mint\_lp

`BurnLp`

burn\_lp

## Event Limits

You can specify 2 types of `limit` on the query:

* `raw`: the max number of events to return.
* `txs`: the max number of transactions to return. **note**: one transaction can emit multiple events, by specifying this limit, you will get all the events associated to the transactions in the response.

## Request

Events by subaccount

Events by product

Events by type

All events

Query events corresponding to specific subaccounts, ordered by `submission index` desc. E.g: all `MatchOrder` events for subaccounts `xxx` specific to spot wBTC.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Query events corresponding to specific products, ordered by `submission index` desc. Uses `txs` limit, will only return a single `tx` and one or more events associated with the `tx`.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Query events corresponding to specific types, ordered by `submission index` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Query all events ordered by `submission index` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

subaccounts

string[]

No

Array of `bytes32` sent as hex strings; each includes the address and the subaccount identifier. When provided, only return events for the specified subaccounts.

product\_ids

number[]

No

when provided, only return events for the specified product ids; return events for all products otherwise.

event\_types

string[]

No

when provided, only return events for the specified event types; return all events otherwise.

idx

number / string

No

when provided, only return events with `submission_idx` <= `idx`

max\_time

number / string

No

when `idx` is not provided, `max_time` (unix epoch in seconds) can be used to only return events created <= `max_time`

limit

object
{"raw": number } or

{"txs": number }

No

* specifying `raw` limit: max number of events to return. defaults to `100`. max possible of `500`.
* specifying `txs` limit: max number of txs to return.

isolated

bool

No

When provided --
- `true`: only returns evens associated to isolated positions.
- `false`: only return events associated to the cross-subaccount.
defaults to `null`. In which case it returns everything.
See [Isolated Margin](https://github.com/nadohq/nado-docs/blob/main/docs/basics/isolated-margin.md) to learn more.

## Response

**Note:**

* the response includes a `txs` field which contains the relevant transactions to the events. There are `>=1 events` per transaction.
* both `events` and `txs` are in descending order by `submission_idx``.`
* use the `submission_idx` to associate an `event` to it's corresponding transaction.

## Response Fields

## Events

* **Net cumulative**: the net difference in that quantity since the beginning of time. For example, if I want to compute total amount paid out in funding between two events, you can subtract the `net_funding_cumulative` of the larger event by the `net_funding_cumulative` of the smaller event.
* **Net unrealized**: similar to `net_cumulative`, but for `net_unrealized`, we have the caveat that when the magnitude of your position decreases, the magnitude of net\_unrealized `decreases` by the same amount.

Field name

Description

submission\_idx

Used to uniquely identify the blockchain transaction that generated the event; you can use it to grab the relevant transaction in the `txs` section.

product\_id

The id of of the product the event is associated with.

event\_type

Name of the transaction type this event corresponds to.

subaccount

The subaccount associated to the event.

pre\_balance

The state of your balance before the event happened.

post\_balance

The state of your balance after the event happened.

product

The state of the product throughout the event.

## Txs

Field name

Description

submission\_idx

Unique identifier of the transaction.

product\_id

Product associated to the transaction.

tx

Raw data of the corresponding transaction e.g: `match_orders` with all associated data.

timestamp

The unix epoch in seconds of when the transaction took place.


Last updated 1 month ago

---


# Fast Withdrawal Signature

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/fast-withdrawal-signature

## Rate limits

* 240 requests/min or 40 requests/10secs per IP address. (**weight = 10**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Fast Withdrawal Signature

Query the signature required for a fast withdrawal at a specific submission index.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "fast_withdrawal_signature": {
    "idx": "12345"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

idx

number / string

Yes

Submission index to fetch the fast withdrawal signature for.

## Response

## Response Fields

## Fast Withdrawal Signature

Field name

Description

signature

Hex string of the signature for fast withdrawal

submission\_idx

Transaction submission index

subaccount

Hex string of the subaccount

product\_id

Product ID (0 for quote asset)

amount

Withdrawal amount (x18 format)

nonce

Nonce for the withdrawal transaction


Last updated 1 month ago

---


# Funding Rate

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/funding-rate

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Single Product

## Request

Funding Rate

Query perp product 24hr funding rate.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "funding_rate": {
    "product_id": 2
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of perp product to fetch funding rate for.

## Response

## Multiple Products

## Request

Perp Prices

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

product\_ids

number[]

Yes

Ids of perp products to fetch funding rate for.

## Response

**Note**: the response is a map of `product_id -> funding_rate` for each requested product.

## Response Fields

Field name

Description

product\_id

Id of the perp product this funding rate corresponds to.

funding\_rate\_x18

Latest 24hr funding rate for the specified product, multiplied by 10^18

update\_time

Epoch time in seconds this funding rate was last updated at


Last updated 1 month ago

---


# Ink Airdrop

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/ink-airdrop

Query the Ink token airdrop allocation for a specific wallet address.

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Ink Airdrop

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "ink_airdrop": {
    "address": "0x1234567890123456789012345678901234567890"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

address

string

Yes

Wallet address (20-byte address) sent as a hex string.

## Response

**Note**: The amount is returned as a string to preserve precision.

## Response Fields

Field name

Description

amount

The Ink token airdrop amount allocated to the address.


Last updated 1 month ago

---


# Interest & funding payments

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/interest-and-funding-payments

## Rate limits

* 480 requests/min or 80 requests/10secs per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Interest and funding

Query subaccount historical interest and funding payments.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "interest_and_funding": {
    "subaccount": "0xD028878bF5c96218E53DA859e587cb8398B17b3f64656661756c740000000000",
    "product_ids": [1, 2],
    "limit": 10,
    "max_idx": 1315836
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccount

string

Yes

A bytes32 sent as a hex string; includes the address and the subaccount identifier.

product\_ids

number[]

Yes

Ids of products to historical interest/funding payments for.

max\_idx

string/number

No

When provided, only return records with `idx` <= `max_idx`.

limit

number

Yes

Max number of records to return. Max possible of `100`.

## Response

## Response Fields

Field name

Description

interest\_payments.product\_id

Id of spot product the interest payment is associated to.

interest\_payments.idx

Id of transaction that triggered the interest payment.

interest\_payments.timestamp

Timestamp of the transaction that triggered the interest payment.

interest\_payments.amount

Amount of interest paid multiplied by 10\*\*18.

interest\_payments.balance\_amount

Previous spot balance at the moment of payment (exclusive of payment amount)

interest\_payments.rate\_x18

Spot interest rate at the moment of payment, multiplied by 10\*\*18.

interest\_payments.oracle\_price\_x18

Oracle price for the spot product at the moment of payment, multiplied by 10\*\*18.

funding\_payments.product\_id

Id of perp product the funding payment is associated to.

funding\_payments.idx

Id of transaction that triggered the funding payment.

funding\_payments.timestamp

Timestamp of the transaction that triggered the funding payment.

funding\_payments.amount

Amount of funding paid multiplied by 10\*\*18.

funding\_payments.balance\_amount

Previous perp balance at the moment of payment +amount of perps locked in LPs (exclusive of payment amount).

funding\_payments.rate\_x18

Perp funding rate at the moment of payment, multiplied by 10\*\*18.

funding\_payments.oracle\_price\_x18

Oracle price for the perp product at the moment of payment, multiplied by 10\*\*18.

next\_idx

Id of the next payment snapshot. Use this as `max_idx` on a subsequent call to get the next page. This will be `null` when there are no more records.


Last updated 1 month ago

---


# Isolated Subaccounts

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/isolated-subaccounts

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

List all isolated subaccounts

List isolated subaccounts for a subaccount

Query all isolated subaccounts.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "isolated_subaccounts": {
    "start_idx": 0,
    "limit": 100
  }
}
```

Query isolated subaccounts associated with a specific subaccount.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "isolated_subaccounts": {
    "subaccount": "0x79cc76364b5fb263a25bd52930e3d9788fcfeea864656661756c740000000000",
    "limit": 100
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccount

string

No

Hex string of the parent subaccount to filter by.

start\_idx

number / string

No

Starting index for pagination. Defaults to 0.

limit

number

No

Max number of isolated subaccounts to return. Defaults to `100`. Max of `500`.

## Response

## Response Fields

## Isolated Subaccounts

Field name

Description

subaccount

Hex string of the parent subaccount

isolated\_subaccount

Hex string of the isolated margin subaccount

product\_id

Product ID for which this isolated subaccount was created

created\_at

Unix epoch time in seconds when the isolated subaccount was created


Last updated 1 month ago

---


# Linked Signer Rate Limit

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/linked-signer-rate-limit

A subaccount can perform a max of 50 [LinkSigner](/developer-resources/api/gateway/executes/link-signer) requests in 7 days. Use this query to check current usage and wait time.

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Link Signer Rate Limit

Queries a subaccount's linked signer rate limits.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "linked_signer_rate_limit": {
    "subaccount": "0x9b9989a4E0b260B84a5f367d636298a8bfFb7a9b42544353504f540000000000"
  }
}
```

## Response

Copy

```
{
  "remaining_tx": "50",
  "wait_time": 0,
  "signer": "0x0000000000000000000000000000000000000000",
  "total_tx_limit": "50"
}
```

**Notes**:

* `remaining_tx`: keeps track of the remaining `LinkSigner` executes that can be performed.
* `total_tx_limit`: that max weekly tx limit.
* `wait_time`: the total seconds you need to wait before performing another `LinkSigner` execute. Can only perform another request when `wait_time` is `0`.
* `signer`: the current linked signer address (20 bytes) associated to the provided `subaccount`. It returns the zero address when no signer is linked.


Last updated 26 days ago

---


# Linked Signers

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/linked-signers

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

List linked signers

Query linked signers ordered by creation time.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "linked_signers": {
    "start_idx": 0,
    "limit": 100
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

start\_idx

number / string

No

Starting index for pagination. Defaults to 0.

limit

number

No

Max number of linked signers to return. Defaults to `100`. Max of `500`.

## Response

## Response Fields

## Linked Signers

Field name

Description

subaccount

Hex string of the subaccount

signer

Hex string of the linked signer address

created\_at

Unix epoch time in seconds when the signer was linked


Last updated 1 month ago

---


# Liquidation Feed

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/liquidation-feed

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Liquidation feed

Queries liquidatable accounts.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "liquidation_feed": {}
}
```

## Response

Copy

```
[
  {
    "subaccount": "0xf2b7cec33cac30582b94979bf03a3cbc73954b2c64656661756c740000000000",
    "update_time": 1680118943
  },
  {
    "subaccount": "0xcb6f1e2ece124a150dcc681c180df2a890432d6a64656661756c740000000000",
    "update_time": 1680118943
  },
  {
    "subaccount": "0x9e6e13be7ea2866c2c7c6e4a118a6c05eee6b44e64656661756c740000000000",
    "update_time": 1680118943
  },
  {
    "subaccount": "0x75008754ffae2889c055961c1b0c5c3ab743c59664656661756c740000000000",
    "update_time": 1680118943
  }
]
```

## Response Fields

Field name

Description

subaccount

Subaccount eligible for liquidation.

update\_time

Last time feed was updated.


Last updated 1 month ago

---


# Market Snapshots

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/market-snapshots

## Rate limits

* IP weight = `max((snapshot_count * product_ids.length / 100), 2)` where `snapshot_count = interval.count.min(500)`. If no `product_ids` are specified, `product_ids.length = 100`.

  + E.g: With `product_ids=[1, 2, 3, 4]` and `interval.count=60`, weight = max((60 \* 4 / 100), 2) = 2, allowing up to 1200 requests per min or 200 requests/10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Market snapshots

Query market snapshots ordered by `timestamp` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
    "market_snapshots": {
        "interval": {
          "count": 2,
          "granularity": 3600,
          "max_time": 1691083697,
        },
        "product_ids": [1, 2]
    }
}
```

## Request Parameters

Parameter

Type

Required

Description

interval

object

Yes

Object to specify desired time period for data

interval.count

number

Yes

Number of snapshots to return, limit 100. Also limited to `interval.count * # product_ids < 2000`

interval.granularity

number

Yes

Granularity value in seconds

interval.max\_time

number / string

No

When providing `max_time` (unix epoch in seconds), only return snapshots with timestamp <= `max_time`. If no value is entered, `max_time` defaults to the current time.

product\_ids

number[]

No

list of product ids to fetch snapshots for, defaults to all products

## Response

**Note**: Please note that this endpoint is currently in beta stage. This feature might be subject to changes without prior notice.

## Response Fields

## Snapshots

**Note**: For product specific fields (i.e. cumulative\_volume, open\_interests), the value is an object which maps product\_ids to their corresponding values.

Field name

Description

timestamp

Timestamp of the snapshot. This may not be perfectly rounded to the granularity since it uses the nearest transaction timestamp less than or equal to `max_time`

cumulative\_users

The cumulative number of subaccounts on Nado. It is updated daily at 9AM ET for historical counts. For current day counts, it is updated every hour.

daily\_active\_users

Daily active users count, updated daily at 9AM ET for historical counts. For current day counts, it is updated every hour.

cumulative\_trades

A map of product\_id -> the cumulative number of trades for the given product\_id.

cumulative\_volumes

A map of product\_id -> cumulative volumes in USDT0 units.

cumulative\_trade\_sizes

A map of product\_id -> cumulative trade sizes in base token

cumulative\_taker\_fees

A map of product\_id -> cumulative taker fees. Taker fees include sequencer fees.

cumulative\_sequencer\_fees

A map of product\_id -> cumulative sequencer fees.

cumulative\_maker\_fees

A map of product\_id -> cumulative maker rebates.

cumulative\_liquidation\_amounts

A map of product\_id -> cumulative liquidation amounts in USDT0 units.

open\_interests

A map of product\_id -> open interests in USDT0 units.

total\_deposits

A map of product\_id -> total deposits held by Nado for a given product at the given time in the base token units.

total\_borrows

A map of product\_id -> total borrows lent by Nado for a given product at the given time in the base token units.

funding\_rates

A map of product\_id -> **hourly** historical funding rates, value returned as **decimal rates** (% = rate \* 100), derived from funding payment amounts. Requires a minimum granularity of 3600 to see non-zero funding rates. Use a granularity where granularity % 3600 = 0 for best results.

deposit\_rates

A map of product\_id -> **daily** deposit rates, values returned as **decimal rates** (% = rate \* 100).

borrow\_rates

A map of product\_id -> **daily** borrow rates, values returned as **decimal rates** (% = rate \* 100).

cumulative\_inflows

A map of product\_id -> cumulative inflows a.k.a deposits in base token units.

cumulative\_outflows

A map of product\_id -> cumulative outflows a.k.a withdraws in base token units.

tvl

The total value locked in USD.


Last updated 1 month ago

---


# Matches

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/matches

## Rate limits

* IP weight = `2 + (limit * subaccounts.length / 10)` where `limit` defaults to 100 (max 500) and `subaccounts.length` defaults to 1

  + E.g: With `limit=100` and 1 subaccount, weight = 12, allowing up to 200 requests per min or 33 requests / 10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Matches by subaccount

Matches by product

Query subaccounts matches ordered by `submission index` desc. Response includes order fill and fee information.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "matches": {
    "product_ids": [
      1,
      2
    ],
    "subaccounts": [
      "0x12a0b4888021576eb10a67616dd3dd3d9ce206b664656661756c740000000000"
    ],
    "max_time": 1679728762,
    "limit": 5,
    "isolated": false
  }
}
```

Query matches for provided products ordered by `submission index` desc. Response includes order fill and fee information.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "matches": {
    "product_ids": [
      1,
      2
    ],
    "max_time": "1679728762",
    "limit": 5
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccounts

string[]

No

Array of `bytes32` sent as hex strings; each includes the address and the subaccount identifier. When provided, only return matches for the specified subaccounts.

product\_ids

number[]

No

When provided, only return matches for the specified product ids; return matches for all products otherwise.

idx

number / string

No

When provided, only return matches with `submission_idx` <= `idx`

max\_time

number / string

No

When `idx` is not provided, `max_time` (unix epoch in seconds) can be used to only return matches created <= `max_time`

limit

number

No

Max number of matches to return. defaults to `100`. max possible of `500`.

isolated

boolean

No

When provided --
- `true`: only returns matches associated to isolated positions.
- `false`: only return matches associated to the cross-subaccount.
defaults to `null`. In which case it returns everything.
See [Isolated Margin](https://github.com/nadohq/nado-docs/blob/main/docs/basics/isolated-margin.md) to learn more.

## Response

**Note:**

* the response includes a `txs` field which contains the relevant transactions for the returned matches. There are `>=1 match events` per transaction.
* both `matches` and `txs` are in descending order by `submission_idx``.`
* use the `submission_idx` to associate a match to it's corresponding transaction.
* the `fee` provided in the response includes taker / maker fees + sequencer fees. See [fees](https://github.com/nadohq/nado-docs/blob/main/docs/basics/fees.md) for more details.

## Response Fields

## Matches

Field name

Description

submission\_idx

Wsed to uniquely identify the blockchain transaction that generated the match; you can use it to grab the relevant transaction in the `txs` section.

isolated

Whether the match is associated with an isolated position. `true` for isolated positions, `false` for cross-subaccount positions.

is\_taker

Whether the order in this match was the taker. `true` if the order was the taker, `false` if the order was the maker.

digest

The unique hash of the order.

order.sender

The sender that placed the order.

order.priceX18

The original order price.

order.amount

The original order amount.

order.expiration

The original order expiration.

order.nonce

The original order nonce.

order.appendix

The original order appendix.

pre\_balance

The state of your balance before the match happened.

post\_balance

The state of your balance after the match happened.

base\_filled

The amount of base (e.g: BTC) filled on this match.

quote\_filled

The amount of quote (e.g: USDT0) filled on this match.

fee

The amount of trading fees + sequencer fees paid on this match.

sequencer\_fee

The amount of sequencer fees paid on this match.

cumulative\_base\_filled

The total amount of base (e.g: BTC) filled on this order up this match.

cumulative\_quote\_filled

The total amount of quote (e.g: USDT0) filled up to this match.

cumulative\_fee

The total amount of fee paid up to this match.

closed\_amount

Total base amount closed by this order (x18).

realized\_pnl

Realized PnL for this order (x18).

## Txs

Field name

Description

submission\_idx

Unique identifier of the transaction.

product\_id

Product associated to the transaction.

taker

The taker order.

maker

The maker order.

timestamp

The unix epoch in seconds of when the transaction took place.


Last updated 16 days ago

---


# NLP Funding Payments

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/nlp-funding-payments

## Rate limits

* 480 requests/min or 80 requests/10secs per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

NLP Funding Payments

Query historical NLP funding payments.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "nlp_funding_payments": {
    "max_idx": "1315836",
    "max_time": "1683315718",
    "limit": 100
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

max\_idx

number / string

No

When provided, only return payments with `idx` <= `max_idx`.

max\_time

number / string

No

When provided, only return payments with `timestamp` <= `max_time` (unix epoch in seconds).

limit

number

No

Max number of payments to return. Defaults to `100`. Max of `500`.

## Response

## Response Fields

## Funding Payments

Field name

Description

product\_id

Id of the perp product

idx

Submission index of the transaction that triggered the payment

timestamp

Unix epoch time in seconds when the payment occurred

total\_payment

Total funding payment amount (x18 format)

rate\_x18

Funding rate used for calculation (x18 format)

oracle\_price\_x18

Oracle price at the time of payment (x18 format)


Last updated 1 month ago

---


# NLP Interest Payments

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/nlp-interest-payments

## Rate limits

* 480 requests/min or 80 requests/10secs per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

NLP Interest Payments

Query historical NLP interest payments.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "nlp_interest_payments": {
    "max_idx": "1315836",
    "max_time": "1683315718",
    "limit": 100
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

max\_idx

number / string

No

When provided, only return payments with `idx` <= `max_idx`.

max\_time

number / string

No

When provided, only return payments with `timestamp` <= `max_time` (unix epoch in seconds).

limit

number

No

Max number of payments to return. Defaults to `100`. Max of `500`.

## Response

## Response Fields

## Interest Payments

Field name

Description

product\_id

Id of the spot product (typically quote/collateral products)

idx

Submission index of the transaction that triggered the payment

timestamp

Unix epoch time in seconds when the payment occurred

amount

Interest payment amount (x18 format)

balance\_amount

Balance amount at the time of payment (x18 format)


Last updated 26 days ago

---


# NLP Snapshots

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/nlp-snapshots

## Rate limits

* Dynamic based on snapshot count (**weight = (limit.min(500) / 100)**)

  + E.g: With `limit=100`, weight = 1
  + E.g: With `limit=500`, weight = 5

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

By interval

By pagination

Query NLP snapshots at specific time intervals.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "nlp_snapshots": {
    "interval": {
      "count": 10,
      "max_time": "1683315718",
      "granularity": 3600
    }
  }
}
```

Query NLP snapshots with pagination.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "nlp_snapshots": {
    "idx": "12345",
    "max_time": "1683315718",
    "limit": 100
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

interval

object

No

Object specifying time interval parameters: `count`, `max_time`, `granularity`

idx

number / string

No

Submission index for pagination.

max\_time

number / string

No

Unix epoch time in seconds. Only return snapshots with timestamp <= `max_time`

limit

number

No

Max number of snapshots to return. Defaults to `100`. Max of `500`.

## Response

## Response Fields

## NLP Snapshots

Field name

Description

submission\_idx

Transaction submission index

timestamp

Unix epoch time in seconds when snapshot was taken

total\_deposits

Total deposits in the NLP pool (x18 format)

total\_borrows

Total borrows from the NLP pool (x18 format)

base\_interest\_rate

Interest rate for base assets (x18 format)

quote\_interest\_rate

Interest rate for quote assets (x18 format)


Last updated 1 month ago

---


# Oracle Price

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/oracle-price

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Oracle Price

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "oracle_price": {
    "product_ids": [1, 2, 3, 4]
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_ids

number[]

Yes

Ids of products to fetch oracles price for.

## Response

## Response Fields

## Prices

Field name

Description

product\_id

Id of product oracle price corresponds to.

oracle\_price\_x18

Latest oracle price multiplied by 10^18.

update\_time

Epoch in seconds the oracle price was last updated at.


Last updated 1 month ago

---


# Oracle Snapshots

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/oracle-snapshots

## Rate limits

* IP weight = `max((snapshot_count * product_ids.length / 100), 2)` where `snapshot_count = interval.count.min(500)`. If no `product_ids` are specified, `product_ids.length = 100`.

  + E.g: With `product_ids=[1, 2, 3, 4]` and `interval.count=60`, weight = max((60 \* 4 / 100), 2) = 2, allowing up to 1200 requests per min or 200 requests/10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Oracle Price

Query oracle snapshots ordered by `timestamp` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
    "oracle_snapshots": {
        "interval": {
          "count": 2,
          "granularity": 3600,
          "max_time": 1691083697,
        },
        "product_ids": [1, 2]
    }
}
```

## Request Parameters

Parameter

Type

Required

Description

interval

object

Yes

Object to specify desired time period for data

interval.count

number

Yes

Number of snapshots to return, limit 100. Also limited to `interval.count * # product_ids < 2000`

interval.granularity

number

Yes

Granularity value in seconds

interval.max\_time

number / string

No

When providing `max_time` (unix epoch in seconds), only return snapshots with timestamp <= `max_time`. If no value is entered, `max_time` defaults to the current time.

product\_ids

number[]

No

list of product ids to fetch snapshots for, defaults to all products

## Response

**Note**: Returns a map of `product_id -> oracle_price`


Last updated 1 month ago

---


# Orders

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/orders

## Rate limits

* IP weight = `2 + (limit * subaccounts.length / 20)` where `limit` defaults to 100 (max 500) and `subaccounts.length` defaults to 1

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Subaccount orders

Orders by digests

Query subaccounts `matched` orders, ordered by `submission index` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "orders": {
    "product_ids": [
      1,
      2
    ],
    "subaccounts": [
      "0x12a0b4888021576eb10a67616dd3dd3d9ce206b664656661756c740000000000"
    ],
    "max_time": 1679728762,
    "trigger_types": [
      "price_trigger",
      "time_trigger"
    ],
    "isolated": false,
    "limit": 5
  }
}
```

Query orders by digests.

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

subaccounts

string[]

conditional

Array of `bytes32` sent as hex strings; each includes the address and the subaccount identifier. Must be provided when querying by `subaccounts`**.**

product\_ids

number[]

No

When provided, only return orders for the specified product ids; return orders for all products otherwise.

idx

number / string

No

When provided, only return orders with `submission_idx` <= `idx`

max\_time

number / string

No

When `idx` is not provided, `max_time` (unix epoch in seconds) can be used to only return orders created <= `max_time`

digests

string[]

conditional

Must be provided when querying by `digests`. only return orders matching the specified digests. **note**: cannot specify digests alongside with `subaccounts` , `product_ids` or `max_time`

trigger\_types

string[]

No

When provided, only return orders matching the specified trigger types. Possible values: `price_trigger`, `time_trigger`, `none`. If not provided, returns orders of all trigger types.

limit

number

No

Max number of orders to return. defaults to `100`. max possible of `500`. **note**: when querying by `digests` limit must be <= total digests provided

isolated

bool

No

When provided --

* `true`: only returns orders associated to isolated positions.
* `false`: only return matches associated to the cross-subaccount.

defaults to `null`. In which case it returns everything.

See [Isolated Margin](https://github.com/nadohq/nado-docs/blob/main/docs/basics/isolated-margin.md) to learn more.

## Response

## Response Fields

Field name

Description

digest

The unique hash of the order.

subaccount

The subaccount that placed the order.

product\_id

The id of of the product the order was executed for.

submission\_idx

Used to uniquely identify the blockchain transaction that generated the order. For multi-fills orders, this is the submission\_idx of the first fill.

last\_fill\_submission\_idx

For multi-fills orders, this is the submission\_idx of the last fill. For single fill orders, it has the same value as `submission_idx`.

amount

The original amount of base to buy or sell.

price\_x18

The original order price.

base\_filled

The total amount of base (e.g: BTC) filled on this order.

quote\_filled

The total amount of quote (e.g: USDT0) filled on this order.

fee

The total amount of fee paid on this order.

closed\_amount

Total base amount closed by this order (x18).

realized\_pnl

Realized PnL from this order (x18).

expiration

The original order expiration.

nonce

The original order nonce.

appendix

The original order appendix.


Last updated 16 days ago

---


# Perp Prices

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/perp-prices

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Single Product

## Request

Perp Prices

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "price": {
    "product_id": 2
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of perp product to fetch prices for.

## Response

## Multiple Products

## Request

Perp Prices

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

product\_ids

number[]

Yes

Ids of perp products to fetch prices for.

## Response

**Note**: the response is a map of `product_id -> perp_prices` for each requested product.

## Response Fields

Field name

Description

product\_id

Id of the perp product.

index\_price\_x18

Latest index price of the perp product, multiplied by 10^18.

mark\_price\_x18

Latest mark price of the perp product, multiplied by 10^18.

update\_time

Epoch time in seconds the perp prices were last updated at.


Last updated 1 month ago

---


# Product Snapshots

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/product-snapshots

## Rate limits

* 240 requests/min or 40 requests/10secs per IP address. (**weight = 10**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Single Product

## Request

Product snapshots

Query snapshots for a given product ordered by `submission index` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
    "products": {
        "product_id": 2,
        "max_time": 1679728762,
        "limit": 1
    }
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

id of product to fetch snapshots for.

idx

number / string

No

when provided, only return product snapshots with `submission_idx` <= `idx`

max\_time

number / string

No

when `idx` is not provided, `max_time` (unix epoch in seconds) can be used to only return snapshots created <= `max_time`

limit

number

No

max number of snapshots to return. defaults to `100`. max possible of `500`.

## Response

**Note**:

* the response includes a `txs` field which contains the relevant transactions to the product snapshots. There are `>=1 product snapshots` per transaction.
* both `products` and `txs` are in descending order by `submission_idx`.
* use the `submission_idx` to associate a `product snapshot` to it's corresponding transaction.

## Response Fields

## Products

Field name

Description

submission\_idx

Used to uniquely identify the blockchain transaction that generated the product snapshot; you can use it to grab the relevant transaction in the `txs` section.

product\_id

The id of of the product the event is associated with.

product

The state of the product at the time of the transaction.

## Txs

Field name

Description

submission\_idx

Unique identifier of the transaction.

tx

Raw data of the corresponding transaction

timestamp

The unix epoch in seconds of when the transaction took place.

## Multiple Products

## Request

Multiple Products snapshots

Query the latest snapshot for the provided products.

`POST [ARCHIVE_ENDPOINT]`

**Body**

## Request Parameters

Parameter

Type

Required

Description

product\_ids

number[]

Yes

Ids of products to fetch snapshots for.

max\_time

number / string

No

When provided, returns the last snapshot created <= `max_time` for each product. Otherwise, the latest snapshot is returned.

## Response

**Note**: the response is a map of `product_id -> snapshot` for each requested product.

## Response Fields

Field name

Description

submission\_idx

Used to uniquely identify the blockchain transaction that generated the product snapshot.

product\_id

The id of of the product the event is associated with.

product

The state of the product at the time of the transaction.


Last updated 1 month ago

---


# Quote Price

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/quote-price

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Get quote price

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "quote_price": {}
}
```

## Response

Copy

```
{
    "price_x18": "999944870000000000"
}
```


Last updated 1 month ago

---


# Sequencer Backlog

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/sequencer-backlog

## Rate limits

* 2400 requests/min or 400 requests/10secs per IP address. (**weight = 1**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Get sequencer backlog

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "backlog": {}
}
```

## Response

Copy

```
{
    "total_txs": "45479039",
    "total_submissions": "45478914",
    "backlog_size": "125",
    "updated_at": "1750365790",
    "backlog_eta_in_seconds": "500",
    "txs_per_second": "0.25"
}
```

## Response Fields

Field name

Description

total\_txs

Total number of transactions stored in the indexer DB.

total\_submissions

Total number of transactions submitted on-chain.

backlog\_size

Number of unprocessed transactions (`total_txs - total_submissions`).

backlog\_eta\_in\_seconds

Estimated time in seconds (`float`) to clear the entire backlog (`null` if unavailable).

txs\_per\_second

Current submission rate in transactions per second (float) (`null` if unavailable).

updated\_at

UNIX timestamp (in seconds) of when the data was last updated.


Last updated 1 month ago

---


# Signatures

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/signatures

## Rate limits

* Dynamic based on `digests` param provided (**weight = 2 + len(digests) / 10**)

  + E.g: With `digests=100`, you can make up to 200 requests per min or 33 requests / 10 secs.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Get order signatures by digests

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "signatures": {
    "digests": [
      "0xf4f7a8767faf0c7f72251a1f9e5da590f708fd9842bf8fcdeacbaa0237958fff",
      "0x0495a88fb3b1c9bed9b643b8e264a391d04cdd48890d81cd7c4006473f28e361"
    ]
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

digests

string[]

Yes

A list of order digests to retrieve signatures for.

## Response

## Response Fields

Field name

Description

digest

The order's generated digest.

signature

The order's generated signature.

signer

The address that signed the order / generated the signature.

is\_linked

Indicates whether this is a signature from a linked signer or the original sender.


Last updated 1 month ago

---


# Subaccount Snapshots

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/subaccount-snapshots

Use this query to get a summary of the latest actions per product on Nado for provided subaccounts. Tracked variables (ex. net interest) are extrapolated to the timestamp or set of timestamps provided.

## Rate limits

* 480 requests/min or 80 requests/10secs per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Subaccount snapshots

Query latest subaccount events/actions ordered by `submission index` desc.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
    "account_snapshots": {
        "subaccounts": [
            "0xec132d41e542c7129268d9d4431f105e0830a81164656661756c745f31000000"
        ],
        "timestamps": [
            1738703761
        ],
        "isolated": false,
        "active": true
    }
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccounts

array

Yes

A list of `bytes32` sent as a hex string; includes the address and the subaccount identifier.

timestamp

array

Yes

A list of timestamps to retrieve multiple subaccount snapshots (one per timestamp).

isolated

boolean

No

A filter to include only isolated or cross margin events.

* If `true`: returns only **isolated** margin events.
* If `false`: returns only **cross** margin events.
* If omitted: returns **both** isolated and cross events.

active

boolean

No

Filters which products to include in the snapshot:

* `true`: returns only products with **non-zero balance** at the timestamp (currently active positions)
* `false`: returns products with **event history** before the timestamp (any historical activity)
* If omitted: defaults to `false`

## Response

Single timestamp

## Response Fields

## Events

* **Net cumulative**: the net difference in that quantity since the beginning of time. For example, if I want to compute total amount paid out in funding between two events, you can subtract the `net_funding_cumulative` of the larger event by the `net_funding_cumulative` of the smaller event.
* **Net unrealized**: similar to `net_cumulative`, but for `net_unrealized`, we have the caveat that when the magnitude of your position decreases, the magnitude of net\_unrealized `decreases` by the same amount.

Field name

Description

submission\_idx

Used to uniquely identify the blockchain transaction that generated the event; you can use it to grab the relevant transaction in the `txs` section.

product\_id

The id of of the product the event is associated with.

event\_type

Name of the transaction type this event corresponds to.

subaccount

The subaccount associated to the event.

pre\_balance

The state of your balance before the event happened.

post\_balance

The state of your balance after the event happened.

product

The state of the product throughout the event.


Last updated 1 month ago

---


# Subaccounts

Source: https://docs.nado.xyz/developer-resources/api/archive-indexer/subaccounts

## Rate limits

* 1200 requests/min or 200 requests/10secs per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

List subaccounts

Find subaccounts by address

Query subaccounts ordered by `subaccount id` ASC.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "subaccounts": {
    "start": 100,
    "limit": 10,
  }
}
```

Query all subaccounts associated to an address ordered by `subaccount id` ASC.

`POST [ARCHIVE_ENDPOINT]`

**Body**

Copy

```
{
  "subaccounts": {
    "address": "0x79CC76364b5Fb263A25bD52930E3d9788fCfEEA8"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

start

string/number

No

Subaccount id to start from (used for pagination). Defaults to 0.

limit

string/number

No

Max number of subaccounts to return. Defaults to 100, max of 500.

address

string

No

An optional wallet address to find all subaccounts associated to it.

## Response

## Response Fields

## Subaccounts

Field name

Description

id

Internal subaccount id

subaccount

Hex string of the subaccount (wallet + subaccount name)

address

Hex string of wallet address

subaccount\_name

Subaccount identifier

created\_at

When subaccount was created

isolated

Whether it's a subaccount for an isolated position


Last updated 1 month ago

---


# Definitions / Formulas

Source: https://docs.nado.xyz/developer-resources/api/definitions-formulas

## Definitions

## **Unsettled USDT0**

Perp balances have two main components:

* `amount`
* `v_quote_balance`

When you buy a perp, `amount` increments and `v_quote_balance` decrements, and vice versa for selling.

Settlement is the process of converting from v\_quote\_balance into actual USDT0 balance. This happens mostly on position close, but may happen on extremely negative PNL positions when we need to pay out positive PNL positions.

The amount that is transferred between `v_quote_balance` in the perp and your USDT0 balance is an amount that results in `amount * oracle_price + v_quote_balance == 0`. Unsettled USDT0 is the total amount that would be transferred between `v_quote_balance` and your USDT0 balance summed across all perps.

## **Unsettled PNL**

**Note:** Technically, there is no such concept as "Unsettled PNL" in our system. However, the UI displays "Unsettled PnL" in some places (e.g., in the USDT0 Balance section) for user clarity.

**What the UI actually shows:** When you see "Unsettled PnL" in the UI, it refers to **Unsettled USDT0** (see above) - the total unsettled quote balance across all perp positions.

**For developers:** Always use **Unsettled USDT0** when referring to this value programmatically. It represents the sum of `amount × oracle_price + v_quote_balance` across all perp positions, which is the amount that would be settled into your USDT0 balance.

## **Unrealized PNL**

Refers to the estimated gains or losses of a current position based on the difference between the average entry price and the current oracle price.

## Formulas

## **Unrealized PNL**

Using the [indexer's events query](/developer-resources/api/archive-indexer/events), your unrealized PNL at the end of some event is given by:

## Total PNL

Your total PNL between `event1` and `event2`, assuming `event1` is after `event2` - is given by:

**Notes**:

* You can use 0 for the second term for the PNL to compute since the beginning of time.
* For spots, we will count deposits and withdraws towards your PNL. i.e. if you deposit BTC, for PNL tracking purposes it is counted as a BTC long at the oracle price.


Last updated 1 month ago

---


# Depositing

Source: https://docs.nado.xyz/developer-resources/api/depositing

There are two ways to deposit funds into Nado programmatically:

1. **On-Chain Contract Call** - Direct interaction with the Endpoint contract (recommended)
2. **Direct Deposit** - Simple transfer to your unique deposit address (alternative, requires caution)

---

## Method 1: On-Chain Contract Call (Recommended)

Deposit by calling the Endpoint contract directly. This is the recommended method for programmatic deposits.

## Contract Address

Find the Endpoint contract address at:

Copy

```
GET <nado-url>/query?type=contracts
```

## Function Interface

## Basic Deposit

Copy

```
function depositCollateral(
    bytes12 subaccountName,  // last 12 bytes of the subaccount bytes32
    uint32 productId,        // product ID for the token
    uint128 amount          // raw token amount (see decimals below)
) external
```

**Parameters:**

* `subaccountName`: The last 12 bytes of your subaccount identifier (e.g., `0x64656661756c740000000000` for "default")
* `productId`: The product ID for the token you're depositing
* `amount`: The raw amount in the token's smallest unit

  + For USDT0 (6 decimals): 1 USDT0 = `1e6` = `1000000`
  + For wETH (18 decimals): 1 wETH = `1e18`
  + For wBTC (8 decimals): 1 wBTC = `1e8`

## Deposit with Referral Code

## Prerequisites

Before depositing via contract call, you must:

1. **Approve Token Allowance**
2. **Get Product Information**

   * Use [All Products](/developer-resources/api/gateway/queries/all-products) query to find:

     + Product ID for your token
     + Token contract address
     + Token decimals

## Example: Depositing 100 USDT0

Assuming USDT0 has product ID `0` and 6 decimals:

## Using Python SDK (Recommended)

## Using Raw Contract (Low-level)

**Advantages:**

* Type-safe: Product ID ensures correct token
* Explicit: Clear which token and amount you're depositing
* Fails early: Transaction reverts if parameters are incorrect

## Processing Time

Deposits may take a few seconds to process after transaction confirmation. You can monitor your balance via:

* [Subaccount Info](/developer-resources/api/gateway/queries/subaccount-info) query
* WebSocket subscriptions for real-time updates

---

## Method 2: Direct Deposit (Alternative)

Each subaccount has a unique deposit address. You can send funds directly to this address and they will automatically be credited to your subaccount.

**Caution**: Only use this method if you're certain you're sending a supported token. Sending unsupported tokens to this address will result in permanent loss of funds. We recommend using the On-Chain Contract Call method instead.

## Getting Your Deposit Address

Query your unique deposit address using the [Direct Deposit Address](/developer-resources/api/archive-indexer/direct-deposit-address) endpoint:

**Request:**

**Response:**

## Depositing Funds

1. Get your deposit address using the API call above
2. **Verify the token is supported** via [All Products](/developer-resources/api/gateway/queries/all-products) query
3. Send the supported token directly to this address from any wallet (MetaMask, CEX, etc.)
4. Funds will be automatically credited to your subaccount within a few seconds

**Advantages:**

* No need to interact with smart contracts
* No need to approve allowances
* Works with any wallet (including CEX withdrawals)
* Simpler for end users

**Critical Warnings:**

* ⚠️ **ONLY send supported tokens** - Find supported tokens via [All Products](/developer-resources/api/gateway/queries/all-products)
* ⚠️ **Sending unsupported tokens will result in permanent loss**
* ⚠️ **Use the correct network** - Ink (mainnet) or Ink Sepolia (testnet)
* ⚠️ **No refunds for incorrect deposits** - Double-check before sending

---

## Important Notes

* **Use Correct Product ID**: Each token has a specific product ID. Using the wrong ID will cause the transaction to fail.
* **Check Token Decimals**: Always multiply by the correct decimal factor (6 for USDT0, 18 for wETH, etc.)
* **Minimum Deposit**: Some products may have minimum deposit amounts
* **Only Supported Tokens**: Only deposit tokens that are listed via the All Products query

---

## Getting Token Information

Use the [All Products](/developer-resources/api/gateway/queries/all-products) query to get:

This information is essential for:

* Finding the correct `productId`
* Getting the token contract for approvals (Method 1 only)
* Calculating the correct `amount` with proper decimals
* Verifying a token is supported before using direct deposit


Last updated 12 days ago

---


# 🔌Endpoints

Source: https://docs.nado.xyz/developer-resources/api/endpoints

## Mainnet

## Ink Mainnet

* **Gateway Websocket**: `wss://gateway.prod.nado.xyz/v1/ws`
* **Gateway REST:** `https://gateway.prod.nado.xyz/v1`
* **Gateway V2:** `https://gateway.prod.nado.xyz/v2`
* **Subscriptions**: `wss://gateway.prod.nado.xyz/v1/subscribe`
* **Archive (Indexer):** `https://archive.prod.nado.xyz/v1`
* **Archive (Indexer) V2:** `https://archive.prod.nado.xyz/v2`
* **Trigger**: `https://trigger.prod.nado.xyz/v1`

## Testnet

## Ink Sepolia

* **Gateway Websocket**: `wss://gateway.test.nado.xyz/v1/ws`
* **Gateway REST:** `https://gateway.test.nado.xyz/v1`
* **Gateway V2:** `https://gateway.test.nado.xyz/v2`
* **Subscriptions**: `wss://gateway.test.nado.xyz/v1/subscribe`
* **Archive (Indexer):** `https://archive.test.nado.xyz/v1`
* **Archive (Indexer) V2:** `https://archive.test.nado.xyz/v2`
* **Trigger**: `https://trigger.test.nado.xyz/v1`


Last updated 1 month ago

---


# Errors

Source: https://docs.nado.xyz/developer-resources/api/errors

List of possible `error` values in the API Response:

## General

Error Code

Error Value

Description

1000/1015

RateLimit

Too Many Requests: You have exceeded the rate limit. Please reduce your request frequency and try again later.

1001

BlacklistedAddress

This address has been blacklisted from accessing the sequencer due to a violation of the Terms of Service. If you believe this is an error, please reach out for assistance.

1002

BlockedLocation

Access from your current location ({location}) is blocked. Please check your location and try again.

1003

BlockedSubdivision

Access from your current location ({location} - {subdivision}) is blocked. Please check your location and try again.

1004

Maintenance

Service is temporarily unavailable due to scheduled maintenance. Please try again later.

## Execute / Query API

Execute Error Response

Copy

```
{
    "status": "failure",
    "signature": {signature},
    "error": "{error msg}",
    "error_code": {error_code}
}
```

Query Error Response

Copy

```
{
    "status": "failure",
    "error": "{error msg}",
    "error_code": {error_code}
}
```

Error Code

Error Value

Description

2000

InvalidPriceIncrement

Invalid order price: Order price, {order.price}, is not divisible by the price\_increment\_x18; price\_increment\_x18 for product {product\_id}: {price\_increment\_x18}.

2001

InvalidAmountIncrement

Invalid order amount: Order amount, {order.amount}, must be divisible by the size\_increment; size\_increment for product {product\_id}: {size\_increment}.

2002

ZeroAmount

Invalid order amount: The provided amount is zero. Please specify a valid order amount.

2003

OrderAmountTooSmall

Invalid order amount: Order amount, {order.amount}, is too small. abs(amount) must be >= min\_size; min\_size for product {product\_id}: {min\_size}.

2004

OrderExpired

Invalid order expiration: The order has already expired. Please ensure the expiration date is in the future.

2005

MaxOrdersLimitReached

You have reached the maximum number of open orders allowed for this market.

2006

UnhealthyOrder

Insufficient account health. The execution of this order would lower your account health below the required threshold. Please adjust your order size or manage your positions to maintain a healthy account balance.

2007

OraclePriceDifference

Order price must be no less than 20% and no more than 500% of the determined oracle price.

2008

PostOnlyOrderCrossesBook

The order cannot be placed as it is post-only and crosses the book. Please adjust your order parameters.

2009

OrderTypeNotSupported

The order type you are trying to use is not currently supported.

2010

InvalidTaker

Invalid taker: The order placement health checks were successfully passed; however, the health checks failed upon matching.

2011

LateRecvExecution

Execute request received after ‘recv\_time’. Ensure that your ‘recv\_time’ allows adequate time for requests to be received.

2012

EarlyRevcExecution

Execute request received more than 100 seconds before the 'recv\_time'. Ensure that the request is sent no more than 100 seconds prior to the 'recv\_time'.

2013

DigestAlreadyExists

The provided digest already exists. Ensure that the provided digest is unique.

2014

UnauthorizedSubaccountCancellation

Operation failed. You're attempting to cancel an order for a different subaccount. Please verify the subaccount.

2015

MarketNotFound

The market for the given product or ticker ID was not found. Please try again with a different product or ticker ID.

2016

InvalidProductId

The provided 'product\_id' is invalid. Please verify and input a valid 'product\_id'.

2017

SpotExecuteExceedsBorrowLimit

Executing this action could result in exceeding your borrowing limit as your spot leverage is currently set to false. Please adjust your withdrawal amount or manage your borrowings to prevent potential risk.

2019

InappropriateSpotLeverage

Spot leverage cannot be applied to a non-spot product. Please ensure you're using the correct type of leverage for the product in question.

2020

OrderNotFound

Order with the provided digest ({digest}) could not be found. Please verify the order digest and try again.

2021

AddressRiskTooHigh

The risk associated with the provided address is too high. Please use a different address or mitigate the associated risk.

2022

InvalidNonce

The provided nonce is invalid. Ensure the nonce is correct and try again.

2023

AddressScreeningPending

Risk screening check for the provided address is still in progress. Please wait until the check is complete before proceeding.

2024

NoPriorDeposit

The provided address has no previous deposits. Ensure you're using an address with prior deposits.

2025

SingleSignatureInsufficientAccountValue

Your account must hold a minimum value of 5 USDT0 to enable single signature sessions. Please ensure your account balance meets this requirement.

2026

DuplicateSignerLinking

You cannot link a signer to the same address more than once. Please provide a unique address for each signer.

2027

SignatureLength

The provided signature does not meet the required length specifications. Please verify and provide a valid signature.

2028

InvalidSigner

The provided signature does not match with the sender's or the linked signer's. Please verify and provide the correct signature.

2029

InvalidSignerZero

Signer cannot be zero. Please provide a valid non-zero signer.

2030

LinkedSignerUpdateLimitExceeded

Linked Signer update limit exceeded. Please wait for {{wait\_time}} seconds before trying again.

2031

FillOrKillNotFilled

Your 'Fill or Kill' order could not be entirely filled. Slippage parameters may be too conservative or size too large.

2033

NonceMissingInPayload

No nonce provided in the request payload. Please ensure a valid nonce is included.

2034

InvalidSignatureV

Invalid Signature: The 'v' value of the signature you provided is not valid. Please verify your signature and try again.

2035

SignatureError

Signature error: {error\_msg}

2036

SubaccountHealthTooLow

Subaccount health insufficient. Please ensure sufficient health level in your subaccount to proceed.

2037

ExcessiveLPTokenBurn

Attempt to burn more LP tokens than currently owned. Please adjust the amount to match or be less than your current LP token balance.

2038

InvalidExecuteMessage

The execute message provided is invalid. Please verify and provide a valid execute message.

2039

MismatchedDigestsAndProductIdsLength

'digests' and 'productIds' arrays should have the same length. Please ensure their lengths match.

2040

InvalidBool

The value you entered is not a valid boolean. Please try again with a value of true or false.

2041

RebateExecuteFormatting

The length of 'subaccounts' array does not match the length of 'amounts' array. Ensure that both arrays have the same number of elements and try again.

2042

NotLiquidatable

Failed to initiate liquidation: The account does not meet the requirements for liquidation.

2043

LiquidatorHealthTooLow

Failed to initiate liquidation: The liquidator's account health is too low.

2044

PositiveInitialHealthLiquidationAttempt

Failed to initiate liquidation: The account to be liquidated has positive initial health.

2045

InvalidLiquidationParameters

Failed to initiate liquidation: Attempted to liquidate quote or provided invalid liquidation parameters.

2046

PerpLiquidationSizeIncrementMismatch

Failed to initiate liquidation: Attempted to liquidate perpetual contract but the amount is not divisible by sizeIncrement.

2047

InvalidLiquidationAmount

Failed to initiate liquidation: Attempted to liquidate either too little, too much or the signs are different.

2048

LiabilitiesBeforePerpsLiquidationAttempt

Failed to initiate liquidation: Attempted to liquidate liabilities before perpetual contracts.

2049

TransferFailed

ERC20 Transfer failed. Please verify the transaction details.

2050

UnauthorizedAction

Unauthorized action attempted. Please ensure you have the necessary permissions.

2051

NotFinalizableSubaccount

Attempted to finalize a subaccount which is not eligible for finalization. Ensure that the subaccount meets all the necessary conditions before proceeding.

2052

InvalidMaker

The maker order subaccount is invalid or has failed the risk check. Please verify the subaccount and ensure it meets the necessary risk parameters.

2053

OrdersCannotBeMatched

Order failed to match due to an internal error. Please try again.

2054

SlippageTooHigh

The requested operation could not be completed due to excessive slippage. Please adjust your order to match market conditions.

2055

InvalidPrice

Invalid price provided. The price must be greater than 0. Please input a valid price.

2056

ImmediateOrCancelDoesNotCross

Your 'Immediate or Cancel' order does not cross the book. Please review the market conditions or adjust your order.

2057

MaxTriggerOrdersLimitReached

You have reached the maximum number of trigger orders allowed for this subaccount.

2058

TriggerOrderNotFound

Trigger order with the provided digest ({digest}) could not be found. Please verify the order digest and try again.

2059

NotTriggerOrder

Submitted order is not a trigger order.

2060

InvalidProductIds

The provided 'product\_ids' is invalid. Please verify input contains only valid products and no duplicates.

2061

InvalidProductType

Invalid product type {{product\_type}}. 'product\_type' must be 'spot' or 'perp'

2062

MissingProductIds

The 'product\_ids' provided is empty. Please ensure you include a non-empty list of valid `product_ids` in your request.

2063

InvalidQueryResponse

Invalid query response. Expected {{Response}}.

2064

ReduceOnlyIncreasesPosition

Reduce only order increases position.

2065

InvalidExpirationBits

Invalid expiration bits: The 4th to 6th most significant bits are reserved and must be unset.

2066

CancelAndPlaceDifferentSenderOrSigner

Sender or signer of cancel and place are not the same.

2067

ReduceOnlyNotTaker

Only taker orders can be set as reduce only.

2068

SystemUnderMaintenance

We're currently performing maintenance on the system. Please try again later.

2069

MarketTradingBlocked

Trading is blocked for this market.

2070

MarketMaxOpenInterest

Market has reached maximum open interest. Please only close positions at this time.

2071

MaxUtilization

Product at maximum utilization

2072

OrderBatchExceedLimit

The number of specified 'orders' exceeds the limit. Please reduce the 'orders' to meet the defined limit.

2073

SelfMatchNotAllowed

Self-match is not allowed.

2074

MismatchedProductIds

Product IDs do not match.

2075

NonDefaultPrivateBatchOrder

Private batch order types must all be default.

2076

InvalidTriggerPrivateBatchOrder

Private batch order cannot be trigger order.

2077

TransferQuoteAmountTooSmall

Transfer quote amount is too small. You must transfer a minimum of 5 USDT0.

2078

TransferQuoteNewRecipientLimitExceeded

Transfer quote to new recipients limit exceeded. Please wait 24hrs before transferring quote to new recipients.

2079

SelfTransferQuoteNotAllowed

Self-transfer quote is not allowed.

2080

WebsocketCompressionRequired

Subscriptions require the header 'Sec-WebSocket-Extensions' with value 'permessage-deflate'.

2081

IsolatedSubaccountCannotPlaceIsolatedOrder

An isolated subaccount cannot place an isolated order.

2082

IsolatedSubaccountInvalidProduct

Invalid product\_id for isolated subaccount.

2083

InvalidIsolatedSpotOrder

Isolated orders cannot be placed on spots.

2084

InvalidIsolatedTriggerOrder

Isolated orders cannot be trigger orders.

2085

InvalidIsolatedReduceOnlyOrder

Isolated orders cannot be reduce-only.

2086

InvalidIsolatedMargin

Isolated margin must be non-negative.

2087

FailedToCreateIsolatedSubaccount

Failed to create isolated subaccount.

2088

InvalidOrderFromIsolatedSubaccount

Orders from isolated subaccount must be reduce-only.

2089

InvalidLinkSignerSender

Cannot link signer to isolated subaccount.

2090

MintNlpAmountTooSmall

Nlp minting amount is too small. You must mint a minimum of 1 USDT0.

2091

AmountTooLarge

Amount is too large.

2092

NAccountHealthTooLow

N\_ACCOUNT health insufficient.

2093

NotCanonicalChain

Can not execute in non-canonical chains.

2094

OrderSizeTooSmall

Invalid order size: Order amount, {order.amount}, or price, {order.price}, is too small. abs(amount) \* price must be >= min\_size; min\_size for product {product\_id}: {min\_size}.

2095

InvalidOrderVersion

Invalid Order Version: the order version in the appendix, {version}, does not match the expected version: {expected\_version}

2096

UnlockedNlpInsufficient

Do not have enough unlocked NLP.

2097

InvalidTriggerOrder

Invalid trigger order appendix.

2098

InvalidTwap

Invalid TWAP order.

2099

InvalidTwapOrderType

TWAP order must be of type 'Immediate or Cancel'. Please ensure your TWAP order uses the correct order type.

2100

InvalidTwapTimes

Invalid TWAP times: {times}. TWAP times must be between 1 and 500.

2101

InvalidTwapAmountDistribution

Invalid TWAP amount distribution: amount {amount} is not evenly divisible by times {times}. For non-random TWAP orders, the total amount must be evenly divisible by the number of executions.

2102

InvalidTwapExpiration

Invalid TWAP expiration: expiration time {expiration} exceeds maximum allowed duration of 25 hours from current time {current\_time}. Please adjust the expiration time.

2103

InvalidTwapInterval

Invalid TWAP interval: interval {interval} seconds exceeds maximum allowed interval of 3600 seconds (1 hour).

2104

InvalidTwapTotalDuration

Invalid TWAP total duration: total duration {duration} seconds exceeds maximum allowed duration of 86400 seconds (24 hours).

2105

InvalidTwapExpirationTiming

Invalid TWAP expiration timing: expiration {expiration} is before the minimum required time {min\_time} for the given interval and times.

2106

InvalidTwapRandomConfiguration

Invalid TWAP random configuration: amounts array presence {amounts\_present} does not match is\_twap\_random flag {is\_random}.

2107

InvalidTwapAmount

Invalid TWAP amount: amount {amount} is zero or has different sign than order amount {order\_amount}.

2108

InvalidTwapAmountsSum

Invalid TWAP amounts sum: sum of amounts {sum} does not match order amount {order\_amount}.

2109

InvalidTwapTriggerAmountConfiguration

Invalid TWAP trigger amount configuration: trigger\_amount presence {trigger\_amount\_present} does not match is\_twap\_random flag {is\_random}.

2110

InvalidTwapIsolated

TWAP orders cannot be isolated. Please remove the isolated flag from your TWAP order.

2111

MaxOrderLimitExceeded

Cannot place more than 50 orders.

2112

NlpPoolAccountsCannotPlaceIsolatedOrder

NLP pool accounts cannot place isolated order.

2113

NlpPoolAccountsCannotPlaceTriggerOrder

NLP pool accounts cannot place trigger order.

2114

BatchSenderMismatch

All orders in a batch must have the same sender. Please ensure all orders are from the same account.

2115

LiquidationFrontrunByNlp

Liquidation succeeded but was executed by the NLP account instead of the requested liquidator.

## Indexer API

Indexer Error Response

Error Code

Error Value

Description

3000

DigestsNotAllowed

Unable to accept 'digests' in conjunction with 'subaccount' or 'product\_ids'. Please make sure your request does not contain these fields simultaneously.

3001

DigestsExceedLimit

The number of specified 'digests' exceeds the specified limit. Please reduce the 'digests' to meet the defined limit.

3002

MissingSubaccount

A 'subaccount' is required but not specified. Please ensure you include a 'subaccount' in your request.

3003

InvalidInterval

Invalid interval: Please try again with a different 'max\_timestamp', 'granularity', or 'count'.

3004

InvalidWithdrawalIdx

Invalid idx: withdrawal tx not found at idx {{idx}}. Check the input idx and try again later.

3005

NotEnoughFastWithdrawalSignatures

Not enough signatures for tx at idx {{idx}}, try again later.

## Others

Error Code

Error Value

Description

4000

PerpTickFormatting

The length of the 'avg\_price\_diffs' array does not match the length of 'product\_ids'. Ensure that the arrays are correctly formed and try again.

4001

NotImplemented

The feature you are trying to use is not yet implemented. Please check back later.

4002

TemporarilyDisabledMintLp

MintLp operation is currently disabled. Please try again later.

4003

EvmRevert

A critical error occurred while attempting match. Reverted with: {revert message}

4004

WithdrawRisk

Protocol risk: {limit} withdrawal limit over 24 hours exceeded; Try again later

5000

InternalError

Internal error: {message}


Last updated 1 month ago

---


# Gateway

Source: https://docs.nado.xyz/developer-resources/api/gateway

There are two types of actions. An `Execute` involves a modification to state, and a `Query` merely fetches information from state.

All actions can be sent over websocket as json payloads at `WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

Additionally, you can send executes and queries over HTTP, at `POST [GATEWAY_REST_ENDPOINT]/execute` and `GET/POST [GATEWAY_REST_ENDPOINT]/query` respectively. For executes, the request should be sent with a json payload, while for queries, the payload should be encoded into url query strings.

`HTTP` requests must set the `Accept-Encoding` to include `gzip`, `br` or `deflate`

## Endpoints

## **Testnet**:

* Websocket: `wss://gateway.test.nado.xyz/v1/ws`
* REST: `https://gateway.test.nado.xyz/v1`

## Websocket

**Notes on** ***keeping websocket connections alive*****:**

* When interacting with our API via websocket, you must send ping frames every 30 seconds to keep the websocket connection alive.
* Ping / Pong frames are built into the websocket protocol and should be supported natively by your websocket library. See [Ping/Pong frames](https://datatracker.ietf.org/doc/html/rfc6455#section-5.5.2) for more info.



Last updated 1 month ago

---


# Executes

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes

## Overview

All executes go through the following endpoint; the exact details of the execution are specified by the JSON payload.

* **Websocket**: `WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`
* **REST**: `POST [GATEWAY_REST_ENDPOINT]/execute`

## **Signing**

All executes are signed using [EIP712](https://eips.ethereum.org/EIPS/eip-712). Each execute request contains:

1. A piece of structured data that includes the sender address
2. A signature of the hash of that structured data, signed by the sender

You can check the SDK for some examples of how to generate these signatures.

See more info in the [signing](/developer-resources/api/gateway/signing) page.

## **Sender Field Structure**

The sender field is a solidity `bytes32` . There are two components:

* an `address` that is a `bytes20`
* a subaccount identifier that is a `bytes12`

For example, if your address was `0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43`, and you wanted to use the default subaccount identifier (i.e: the word `default`) you can set `sender` to `0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c4364656661756c740000000000` , which sets the subaccount identifier to `64656661756c740000000000`.

## **Amounts**

For `DepositCollateral` and `WithdrawCollateral`, the amount specifies the physical token amount that you want to receive. `i.e.` if USDT0 has 6 decimals, and you want to deposit or withdraw 1 USDT0, you specify `amount = 1e6`.

For all other transactions, amount is normalized to 18 decimals, so `1e18` == one unit of the underlying asset. For example, if you want to buy 1 wETH, regardless of the amount of decimals the wETH contract has on chain, you specify `1e18` in the amount field of the order.

## API Response

All `Execute` messages return the following information:

## Success

## Failure


Last updated 1 month ago

---


# Burn NLP

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/burn-nlp

## Rate limits

* 60 burns/min or 10 burns every 10 seconds per wallet. (**weight = 10**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "burn_nlp": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "nlpAmount": "10001000000000000000000"
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "burn_lp": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productId": 1,
      "amount": "10001000000000000000000"
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Burn NLP transaction object. See [Signing](/developer-resources/api/gateway/executes/burn-nlp#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.nlpAmount

string

Yes

Amount of NLP tokens to burn multiplied by 1e18, sent as a string.

tx.nonce

string

Yes

This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/gateway/executes/burn-nlp#signing) section for more details.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`nlpAmount`: amount of NLP tokens to burn, sent as a string. This must be positive and must be specified with 18 decimals.

`nonce`: the `tx_nonce`. This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Cancel And Place

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/cancel-and-place

## Rate limits

* The sum of [Cancel Orders](/developer-resources/api/gateway/executes/cancel-orders#rate-limits) + [Place Order](/developer-resources/api/gateway/executes/place-order#rate-limits) limits

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "cancel_and_place": {
    "cancel_tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productIds": [2],
      "digests": ["0x"],
      "nonce": "1"
    },
    "cancel_signature": "0x",
    "place_order": {
      "product_id": 1,
      "order": {
        "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
        "priceX18": "1000000000000000000",
        "amount": "1000000000000000000",
        "expiration": "4294967295",
        "appendix": "1537",
        "nonce": "1757062078359666688"
      },
      "signature": "0x",
    }
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

## Request Parameters

Parameter

Type

Required

Description

cancel\_tx

object

Yes

Cancel order transaction object. See [Cancel order signing](/developer-resources/api/gateway/executes/cancel-orders#signing) for details on the transaction fields.

cancel\_tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

cancel\_tx.productIds

number[]

Yes

A list of product IDs, corresponding to the product ids of the orders in `digests`

cancel\_tx.digests

string[]

Yes

A list of order digests, represented as hex strings.

cancel\_tx.nonce

string

Yes

Used to differentiate between the same cancellation multiple times. See [Cancel order signing](/developer-resources/api/gateway/executes/cancel-orders#signing) section for more details.

cancel\_signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/gateway/executes/cancel-and-place#signing)[Cancel order signing](/developer-resources/api/gateway/executes/cancel-orders#signing) for more details.

place\_order

object

Yes

Payload of order to be placed. See [Place order request parameters](/developer-resources/api/trigger/executes/place-order#request-parameters) for payload details.

## Signing

**Note**: both `cancel_tx` and `place_order` objects must be signed using the same signer, otherwise the request will be rejected.

* See [Cancel orders signing](/developer-resources/api/gateway/executes/cancel-orders#signing) for details on how to sign the order cancellation.
* See [Place order signing](/developer-resources/api/gateway/executes/place-order#signing) for details on how to sign the order placement.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Cancel Orders

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/cancel-orders

## Rate limits

* When no **digests** are provided: 600 cancellations/min or 10 cancellations/sec per wallet. (**weight=1**)
* When **digests** are provided: 600/(total digests) cancellations per minute per wallet. (**weight=total digests**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "cancel_orders": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productIds": [2],
      "digests": ["0x"],
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "cancel_orders": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productIds": [0],
      "digests": ["0x"],
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Cancel order transaction object. See [Signing](/developer-resources/api/gateway/executes/cancel-orders#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.productIds

number[]

Yes

A list of product IDs, corresponding to the product ids of the orders in `digests`

tx.digests

string[]

Yes

A list of order digests, represented as hex strings.

tx.nonce

string

Yes

Used to differentiate between the same cancellation multiple times. See [Signing](/developer-resources/api/gateway/executes/cancel-orders#signing) section for more details.

signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/gateway/executes/cancel-orders#signing) section for more details.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier

`productIds`: a list of product IDs, corresponding to the product ids of the orders in `digests`

`digests`: a list of order digests, represented as hex strings, for the orders you want to cancel.

`nonce`: used to differentiate between the same cancellation multiple times, and a user trying to place a cancellation with the same parameters twice. Sent as a string. Encodes two bit of information:

* Most significant `44` bits encoding the `recv_time` in milliseconds after which the cancellation should be ignored by the matching engine; the engine will accept cancellations where `current_time < recv_time <= current_time + 100000`
* Least significant `20` bits are a random integer used to avoid hash collisions

  For example, to place a cancellation with a random integer of `1000`, and a discard time 50 ms from now, we would send a nonce of `(timestamp_ms() + 50) << 20 + 1000`

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Cancel Product Orders

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/cancel-product-orders

## Rate limits

* When no **productIds** are provided**:** 12 cancellations/min or 2 cancellations/sec per wallet. (**weight=50**)
* When **productIds** are provided: 600 / (5 \* total productIds) cancellations per minute per wallet. (**weight=5\*total productIds**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "cancel_product_orders": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productIds": [2],
      "nonce": "1"
    },
    "signature": "0x",
    "digest": null
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "cancel_product_orders": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productIds": [0],
      "nonce": "1"
    },
    "signature": "0x",
    "digest": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Cancel product orders transaction object. See [Signing](/developer-resources/api/gateway/executes/cancel-product-orders#signing) section for details on transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.productIds

number[]

Yes

A list of product IDs to cancel orders for.

tx.nonce

string

Yes

Used to differentiate between the same cancellation multiple times. See [Signing](/developer-resources/api/gateway/executes/cancel-product-orders#signing) section for more details.

signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/gateway/executes/cancel-product-orders#signing) section for more details.

digest

string

No

Hex string representing a hash of the `CancellationProducts` object.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier

`productIds`: a list of product Ids for which to cancel all subaccount orders. When left empty, orders from all products will be cancelled.

`nonce`: used to differentiate between the same cancellation multiple times, and a user trying to place a cancellation with the same parameters twice. Sent as a string. Encodes two bit of information:

* Most significant `44` bits encoding the `recv_time` in milliseconds after which the cancellation should be ignored by the matching engine; the engine will accept cancellations where `current_time < recv_time <= current_time + 100000`
* Least significant `20` bits are a random integer used to avoid hash collisions

  For example, to place a cancellation with a random integer of `1000`, and a discard time 50 ms from now, we would send a nonce of `(timestamp_ms() + 50) << 20 + 1000`

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Link Signer

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/link-signer

Each subaccount can have at most one linked signer at a time. A linked signer can perform any execute on behalf of the subaccount it is linked to. Use the [Linked Signer](/developer-resources/api/gateway/queries/linked-signer) query to view your current linked signer.

**Please note**:

* To enable a linked signer, your subaccount must have a minimum of **5 USDT0** worth in account value.

## Rate limits

* A max of 50 link signer requests every 7 days per subaccount. (**weight=30**). Use the [Linked Signer Rate Limit](/developer-resources/api/archive-indexer/linked-signer-rate-limit) query to check a subaccount's linked signer usage and remaining wait time.

See more general details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "link_signer": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "signer": "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "link_signer": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "signer": "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

A link signer transaction object. See [Signing](/developer-resources/api/gateway/executes/link-signer#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.signer

string

Yes

A `bytes32` sent as a hex string; includes the address (first 20 bytes) that'll be used as the `sender's` signer. the last 12 bytes can be set to anything.

tx.nonce

string

Yes

This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/gateway/executes/link-signer#signing) section for more details.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier of the primary subaccount to add a signer to.

`signer`: a `bytes32` sent as a hex string; includes the address (first 20 bytes) that'll be used as the `sender's` signer.

**Notes**:

* the last 12 bytes of the `signer` field do not matter and can be set to anything.
* set `signer` to the zero address to revoke current signer on the provided `sender`.

`nonce`: the `tx_nonce`. This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 26 days ago

---


# Liquidate Subaccount

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/liquidate-subaccount

## Rate limits

* 30 liquidations/min or 5 liquidations every 10 seconds per wallet. (**weight=20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "liquidate_subaccount": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "liquidatee": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productId": 1,
      "isEncodedSpread": false,
      "amount": "1000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "liquidate_subaccount": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "liquidatee": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "mode": 0,
      "healthGroup": 1,
      "amount": "1000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Liquidate subaccount transaction object. See [Signing](/developer-resources/api/gateway/executes/liquidate-subaccount#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.liquidatee

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the subaccount being liquidated.

tx.productId

number

Yes

Perp Liquidation:

* A valid perp product Id.

Spot Liquidation:

* A valid spot product Id.

Spread Liquidation:

* An encoded perp / spot product Ids, where the lower 16 bits represent the spot product and the higher 16 bits represent the perp product. `isEncodedSpread` must be set to `true` for spread liquidation. See [Signing](/developer-resources/api/gateway/executes/liquidate-subaccount#signing) section for more details.

tx.isEncodedSpread

bool

Yes

When set to `true`, the `productId` is expected to encode a perp and spot product Ids as follows: `(perp_id << 16) | spot_id`

tx.amount

string

Yes

The amount to liquidate multiplied by 1e18, sent as a string.

tx.nonce

string

Yes

This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/gateway/executes/liquidate-subaccount#signing) section for more details.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`liquidatee`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`productId`: The product to liquidate as well as the liquidation mode.

* *Perp liquidation* ⇒ A valid `perp` product id is provided and `isEncodedSpread` is set to `false`.
* *Spot liquidation* ⇒ A valid `spot` product id is provided and `isEncodedSpread` is set to `false`
* *Spread Liquidation* => If there are perp and spot positions in different directions, liquidate both at the same time. Must be set to a 32 bits integer where the lower 16 bits represent the `spot` product and the higher 16 bits represent the `perp` product. `isEncodedSpread` must be set to `true`.

***Computing\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\**** ***productId*** ***\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*for Spread Liquidation***

`isEncodedSpread`: indicates whether `productId` encodes both a `spot` and a `perp` product Id for spread liquidation.

`amount`: the amount to liquidate multiplied by 1e18, sent as a string. Can be positive or negative, depending on if the user’s balance is positive or negative.

`nonce`: the `tx_nonce`. This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Mint NLP

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/mint-nlp

## Rate limits

* Wallet weight = `10` - allows 60 mints/min or 10 mints every 10 seconds per wallet.

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "mint_nlp": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "quoteAmount": "1000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "mint_nlp": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productId": 1,
      "amountBase": "1000000000000000000",
      "quoteAmountLow": "10000000000000000000000",
      "quoteAmountHigh": "20000000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Mint NLP transaction object. See [Signing](/developer-resources/api/gateway/executes/mint-nlp#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.quoteAmount

string

Yes

This amount of quote to be consumed by minting NLPs multiplied by 1e18, sent as a string.

tx.nonce

string

Yes

This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

signature

string

Yes

Hex string representing hash of the **signed** transaction. See [Signing](/developer-resources/api/gateway/executes/mint-nlp#signing) section for more details.

spot\_leverage

boolean

No

Indicates whether leverage should be used; when set to `false` , the mint fails if the transaction causes a borrow on the subaccount. Defaults to `true`.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`quoteAmount`: this is the amount of quote to be consumed by minting NLPs, sent as a string. This must be positive and must be specified with 18 decimals.

`nonce`: the `tx_nonce`. This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 11 days ago

---


# Place Order

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/place-order

## Rate limits

* With spot leverage: 600 orders/minute or 10 orders/sec per wallet. (**weight=1**)
* Without spot leverage: 30 orders/min or 5 orders every 10 seconds per wallet. (**weight = 20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "place_order": {
    "product_id": 1,
    "order": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "priceX18": "1000000000000000000",
      "amount": "1000000000000000000",
      "expiration": "4294967295",
      "nonce": "1757062078359666688",
      "appendix": "1"
    },
    "signature": "0x",
    "id": 100
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "place_order": {
    "product_id": 1,
    "order": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "priceX18": "1000000000000000000",
      "amount": "1000000000000000000",
      "expiration": "4294967295",
      "nonce": "1757062078359666688"
    },
    "signature": "0x",
    "id": 100
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of spot / perp product for which to place order. Use [All products](/developer-resources/api/gateway/queries/all-products) query to retrieve all valid product ids.

order

object

Yes

Order object, see [Signing](/developer-resources/api/gateway/executes/place-order#signing) section for details on each order field.

order.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

order.priceX18

string

Yes

Price of the order multiplied by 1e18.

order.amount

string

Yes

Quantity of the order multiplied by 1e18.

order.expiration

string

Yes

A time after which the order should automatically be cancelled, as a timestamp in seconds after the unix epoch.

order.nonce

string

Yes

Used to differentiate between the same order multiple times. See [Signing](/developer-resources/api/gateway/executes/place-order#signing) section for more details.

order.appendix

string

Yes

Encodes various order properties including execution types, isolated positions, TWAP parameters, and trigger types. See order appendix section for more details.

signature

string

Yes

Hex string representing hash of the **signed** order. See [Signing](/developer-resources/api/gateway/executes/place-order#signing) section for more details.

digest

string

No

Hex string representing a hash of the order.

spot\_leverage

boolean

No

Indicates whether leverage should be used; when set to `false` , placing the order fails if the transaction causes a borrow on the subaccount. Defaults to `true`.

id

number

No

An optional id that when provided is returned as part of `Fill` and `OrderUpdate` stream events. See [subscriptions](/developer-resources/api/subscriptions) for more details.
**NOTE**: The client `id` should not be used to differentiate orders, as it is not included in the order hash (i.e., the order `digest`). Instead, use the last 20 bits of the order nonce to distinguish between similar orders. For more details, refer to [Order Nonce](/developer-resources/api/gateway/executes/place-order#order-nonce).

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier

`priceX18`: an `int128` representing the price of the order multiplied by 1e18, sent as a string. For example, a price of 1 USDT0 would be sent as `"1000000000000000000"`

`amount`: an `int128` representing the quantity of the order multiplied by 1e18, sent as a string. A positive amount means that this is a buy order, and a negative amount means this is a sell order.

`expiration`: a time after which the order should automatically be cancelled, as a timestamp in seconds after the unix epoch, sent as a string.

## Order Nonce

`nonce`: used to differentiate between the same order multiple times, and a user trying to place an order with the same parameters twice. Sent as a string. Encodes two bit of information:

* Most significant `44` bits encoding the time in milliseconds (a `recv_time`) after which the order should be ignored by the matching engine
* Least significant `20` bits are a random integer used to avoid hash collisions

  For example, to place an order with a random integer of `1000`, and a discard time 50 ms from now, we would send a nonce of `((timestamp_ms() + 50) << 20) + 1000)`

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Order Appendix

See more details and examples in our [Order Appendix](/developer-resources/api/order-appendix) page.

`appendix`: is a 128-bit integer that encodes extra order parameters like execution type, isolated margin, and trigger type.

## Bit Layout

**Fields (from LSB to MSB):**

* **Version (8 bits, 0–7)** – protocol version (currently `1`)
* **Isolated (1 bit, 8)** – whether the order uses isolated margin
* **Order Type (2 bits, 9–10)** – 0 = DEFAULT, 1 = IOC, 2 = FOK, 3 = POST\_ONLY

  + `0` - `DEFAULT`: Standard limit order behavior
  + `1` - `IOC (Immediate or Cancel)`: Execute immediately, cancel unfilled portion
  + `2` - `FOK (Fill or Kill)`: Execute completely or cancel entire order
  + `3` - `POST_ONLY`: Only add liquidity, reject if would take liquidity
* **Reduce Only (1 bit, 11)** – only decreases an existing position.
* **Trigger Type (2 bits, 12–13)** – 0 = NONE, 1 = PRICE, 2 = TWAP, 3 = TWAP\_CUSTOM\_AMOUNTS
* **Reserved (50 bits, 14–63)** – future use
* **Value (64 bits, 64–127)** – extra data (isolated margin or TWAP parameters)

  + if `trigger` is `2` or `3` ⇒ `value` represents how many times the TWAP order will execute and the maximum acceptable slippage. Encoded as:

    - `times` : Number of TWAP executions.
    - `slippage_x6`: Maximum slippage × 1,000,000 (6 decimal precision).
  + if `isolated` is `1` ⇒ `value` represents `margin_x6` (in x6 precision, 6 decimals) to be transferred to the isolated subaccount when the order gets its first match.
  + otherwise, `value` is `0`.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Place Orders

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/place-orders

Place multiple orders in a single request. This is more efficient than placing orders individually and allows for better control over batch order placement.

## Rate limits

* With spot leverage: 600 orders/minute or 10 orders/sec per wallet. (**weight=1 per order**)
* Without spot leverage: 30 orders/min or 5 orders every 10 seconds per wallet. (**weight = 20 per order**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

**Note**: There is a 50ms processing penalty for each `place_orders` request to ensure fair sequencing and prevent gaming of the matching engine.

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "place_orders": {
    "orders": [
      {
        "product_id": 2,
        "order": {
          "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
          "priceX18": "100000000000000000000000",
          "amount": "1000000000000000000",
          "expiration": "4294967295",
          "nonce": "1757062078359666688",
          "appendix": "1"
        },
        "signature": "0x...",
        "id": 100
      },
      {
        "product_id": 3,
        "order": {
          "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
          "priceX18": "3800000000000000000000",
          "amount": "2000000000000000000",
          "expiration": "4294967295",
          "nonce": "1757062078359666689",
          "appendix": "1"
        },
        "signature": "0x...",
        "id": 101
      }
    ],
    "stop_on_failure": false
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

## Request Parameters

Parameter

Type

Required

Description

orders

array

Yes

Array of order objects to place. Each order follows the same structure as [Place Order](/developer-resources/api/gateway/executes/place-order).

orders[].product\_id

number

Yes

Id of spot / perp product for which to place order.

orders[].order

object

Yes

Order object (same structure as single order placement).

orders[].signature

string

Yes

Hex string representing hash of the **signed** order.

orders[].digest

string

No

Hex string representing a hash of the order.

orders[].spot\_leverage

boolean

No

Indicates whether leverage should be used for this order. Defaults to `true`.

orders[].id

number

No

An optional id returned in `Fill` and `OrderUpdate` events.

stop\_on\_failure

boolean

No

If `true`, stops processing remaining orders when the first order fails. Already successfully placed orders are NOT cancelled. Defaults to `false`.

## Response

## Response Fields

Field

Description

digest

Order digest (32-byte hash) if successfully placed, `null` if failed.

error

Error message if order failed, `null` if successful.

## Behavior

* **Partial Success**: By default, orders are processed independently. Some orders may succeed while others fail.
* **Stop on Failure**: Set `stop_on_failure: true` to stop processing remaining orders when the first order fails. Already successfully placed orders remain on the book.
* **Order Signing**: Each order must be individually signed using EIP712 (see [Signing](/developer-resources/api/gateway/signing) for details).
* **Rate Limits**: Rate limit weight is calculated per order (1 per order with leverage, 20 per order without).

## Use Cases

* **Spread Trading**: Place both legs of a spread trade in one request
* **Multiple Markets**: Open positions across multiple products in one request

## Example

Placing BTC and ETH perp orders simultaneously:

## See Also

* [Place Order](/developer-resources/api/gateway/executes/place-order) - Single order placement
* [Cancel And Place](/developer-resources/api/gateway/executes/cancel-and-place) - Atomic cancel and place
* [Signing](/developer-resources/api/gateway/signing) - EIP712 order signing


Last updated 1 month ago

---


# Transfer Quote

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/transfer-quote

## Fees

Transfers between subaccounts incur a network fee:

* **Standard transfers**: 1 USDT0
* **Isolated subaccount transfers**: 0.1 USDT0 (when either sender or recipient is an isolated subaccount)

The fee is automatically deducted from the sender's balance.

## Rate limits

* 60 transfer quotes/min or 10 every 10 seconds per wallet. (**weight=10**)
* A max of 5 transfer quotes to new recipients (subaccounts) every 24hrs.

  + **Note**: Transferring quote to a subaccount that doesn't exist, creates the subaccount.

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "transfer_quote": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "recipient": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743100000000000000",
      "amount": "10000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "transfer_quote": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "recipient": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743100000000000000",
      "amount": "10000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Transfer Quote transaction object. See [Signing](/developer-resources/api/gateway/executes/transfer-quote#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.recipient

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the quote recipient.

tx.amount

string

Yes

The amount of USDT0 to transfer, denominated in `x18`. Transfr amount must be `>= 5 USDT0` . See [Signing](/developer-resources/api/gateway/executes/transfer-quote#signing) section for more details.

tx.nonce

string

Yes

This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

signature

string

Yes

Hex string representing hash of the **signed** transaction. See [Signing](/developer-resources/api/gateway/executes/transfer-quote#signing) section for more details.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`recipient`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`amount`: the amount of quote to transfer, sent as an `x18` string.

**Notes:**

* If you are transferring `5 USDT0`, must specify `5000000000000000000` i.e 5 USDT0 \* 1e18.
* Transfer amount should be >= 5 USDT0.

`nonce`: the `tx_nonce`. This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 26 days ago

---


# Withdraw Collateral

Source: https://docs.nado.xyz/developer-resources/api/gateway/executes/withdraw-collateral

**Note**: use the [max withdrawable](/developer-resources/api/gateway/queries/max-withdrawable) query to determine the max amount you can withdraw for a given spot product.

## Rate limits

* With spot leverage: 60 withdrawals/min or 10 withdrawals every 10 seconds per wallet. (**weight = 10**)
* Without spot leverage: 30 withdrawals/min or 5 withdrawals every 10 seconds per wallet. (**weight=20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits).

## Request

Websocket

REST

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "withdraw_collateral": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productId": 1,
      "amount": "1000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

`POST [GATEWAY_REST_ENDPOINT]/execute`

**Body**

Copy

```
{
  "withdraw_collateral": {
    "tx": {
      "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
      "productId": 1,
      "amount": "1000000000000000000",
      "nonce": "1"
    },
    "signature": "0x"
  }
}
```

## Request Parameters

Parameter

Type

Required

Description

tx

object

Yes

Withdraw collateral transaction object. See [Signing](/developer-resources/api/gateway/executes/withdraw-collateral#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.productId

number

Yes

A spot product ID to withdraw from.

tx.amount

string

Yes

The amount of the asset to withdraw, denominated in the base ERC20 token of the specified product e.g: USDT0 (product=0) has 6 decimals whereas wETH (product=3) has 18. See [Signing](/developer-resources/api/gateway/executes/withdraw-collateral#signing) section for more details.

tx.nonce

string

Yes

This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

signature

string

Yes

Hex string representing hash of the **signed** transaction. See [Signing](/developer-resources/api/gateway/executes/withdraw-collateral#signing) section for more details.

spot\_leverage

boolean

No

Indicates whether leverage should be used; when set to `false` , the withdrawal fails if the transaction causes a borrow on the subaccount. Defaults to `true`.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier.

`productId`: a `uint32` that specifies the product you’d like to withdraw collateral from; must be for a spot product.

`amount`: the amount of asset to withdraw, sent as a string. Note that this is different from the amounts provided in transactions that aren’t `depositCollateral`. This is the raw amount of the ERC20 token you want to receive, i.e. if USDT0 has 6 decimals and you want to withdraw 1 USDT0, specify 1e6; if wETH has 18 decimals and you want to withdraw 1 wETH, specify 1e18. Use [all products](/developer-resources/api/gateway/queries/all-products) query to view the token address of the corresponding product which can be used to determine the correct decimals to use.

`nonce`: the `tx_nonce`. This is an incrementing nonce, can be obtained using the [Nonces](/developer-resources/api/gateway/queries/nonces) query.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `nonce` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

## Failure


Last updated 1 month ago

---


# Queries

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries

All queries go through the following endpoint; the exact details of the query are specified by query params or `Websocket` messages.

* **Websocket**: `WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`
* **REST**: `GET [GATEWAY_REST_ENDPOINT]/query` or `POST [GATEWAY_REST_ENDPOINT]/query`

## Overview

## **Amounts and Prices**

In general, amounts come back normalized to 18 decimal places. Meaning that for a balance of 1 USDT0, regardless of the number of decimals USDT0 has on-chain, a value of 1e18 will be returned.

Prices are in `x18`, so if the price of one wBTC is $20,000, regardless of the number of decimals wBTC has on-chain, the price will be returned as `20,000 * 1e18`.

## API Response

All `queries` return in the format:

Copy

```
{
  "status": "success" | "failure",
  "data"?: {data},
  "error"?: "{error_msg}",
  "error_code"?: {error_code},
  "request_type": "{request_type}"
}
```


Last updated 1 month ago

---


# All Products

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/all-products

## Rate limits

* 480 requests/min or 8 requests/sec per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
 "type": "all_products"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=all_products`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
 "type": "all_products"
}
```

## Response

**Note**:

* A product is some asset / position an account can take on.
* A market is a venue for a product against USDT0.
* All products have a market quoted against USDT0, except for product 0.
* Product 0 is the USDT0 asset itself.
* You can retrieve product symbols via [Symbols](/developer-resources/api/symbols) query.Body


Last updated 1 month ago

---


# Edge All Products

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/edge-all-products

## Rate limits

* 480 requests/min or 8 requests/sec per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
 "type": "edge_all_products"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=edge_all_products`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
 "type": "edge_all_products"
}
```

## Response

**Note**:

* A product is some asset / position an account can take on.
* A market is a venue for a product against USDT0.
* All products have a market quoted against USDT0, except for product 0.
* Product 0 is the USDT0 asset itself.
* You can retrieve product symbols via [Symbols](/developer-resources/api/symbols) query.Body
* Returns a mapping of `chain_id -> all_products`


Last updated 1 month ago

---


# Fee Rates

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/fee-rates

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "fee_rates",
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=fee_rates&sender={sender}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
  "type": "fee_rates",
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000"
}
```

## Request Parameters

Parameter

Type

Required

Description

sender

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier.

## Response

* `taker_fee_rates_x18`: taker fee associated with a given product indexed by `product_id`. **Note**: this fee represents the basis point (BPS) on a taker order in `x18`.
* `maker_fee_rates_x18`: maker fee associated with a given produced indexed by `product_id``.`
* `withdraw_sequencer_fees`: withdraw fees associated with a given product indexed by `product_id`. **Note**: this fee represents a fixed amount of product to be deducted as fee in `x18`.

See our [fees](https://github.com/nadohq/nado-docs/blob/main/docs/basics/fees.md) page for details about current fee rates.


Last updated 1 month ago

---


# Health Groups

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/health-groups

**Note**: a health group is a perp and spot product whose health is calculated together (e.g. BTC and BTC-PERP).

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "health_groups"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=health_groups`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
  "type": "health_groups"
}
```

## Response

Copy

```
{
    "status": "success",
    "data": {
        "health_groups": [
            [
                1,
                2
            ]
        ]
    },
    "request_type": "query_health_groups"
}
```

* `health_groups`: list of all available health groups. **Note**: `health_groups[i]` is the spot / perp product pair of health group `i` where `health_groups[i][0]` is the spot `product_id` and `health_groups[i][1]` is the perp `product_id`. Additionally, it is possible for a health group to only have either a spot or perp product, in which case, the product that doesn’t exist is set to `0`.


Last updated 1 month ago

---


# Insurance

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/insurance

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "insurance"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=insurance`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
  "type": "insurance"
}
```

## Response

Copy

```
{
    "status": "success",
    "data": {
        "insurance": "552843342443351553629462"
    },
    "request_type": "query_insurance"
}
```


Last updated 1 month ago

---


# Isolated Positions

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/isolated-positions

## Rate limits

* 240 requests/min or 40 requests every 10 seconds per IP address. (**weight = 10**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
    "type": "isolated_positions",
    "subaccount": "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=isolated_positions&subaccount={subaccount}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
    "type": "isolated_positions",
    "subaccount": "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000"
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccount

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier. See [sender field structure](/developer-resources/api/gateway/executes#sender-field-structure) for details.

## Response

**Note**:

* `isolated_positions[i].subaccount`: is the isolated subaccount for the base product.
* `healths`:

  + `healths[0]`: info about your initial health, which is weighted by `long_weight_initial_x18` and `short_weight_initial_x18.`
  + `healths[1]`: info about your maintenance health, which is weighted by `long_weight_maintenance_x18` and `short_weight_maintenance_x18.`
  + `healths[2]`: info about your unweighted health.


Last updated 1 month ago

---


# Linked Signer

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/linked-signer

## Rate limits

* 480 requests/min or 8 requests/sec per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "linked_signer",
  "subaccount": "0x9b9989a4E0b260B84a5f367d636298a8bfFb7a9b42544353504f540000000000"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=linked_signer&subaccount=0x9b9989a4E0b260B84a5f367d636298a8bfFb7a9b42544353504f540000000000`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
  "type": "linked_signer",
  "subaccount": "0x9b9989a4E0b260B84a5f367d636298a8bfFb7a9b42544353504f540000000000"
}
```

## Response

Copy

```
{
  "status": "success",
  "data": {
    "linked_signer": "0x0000000000000000000000000000000000000000"
  },
  "request_type": "query_linked_signer",
}
```

**Notes**:

* `linked_signer`: the current linked signer address (20 bytes) associated to the provided `subaccount`. It returns the zero address when no signer is linked.


Last updated 1 month ago

---


# Market Liquidity

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/market-liquidity

## Rate limits

* 2400 requests/min or 40 requests/sec per IP address. (**weight = 1**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "market_liquidity",
  "product_id": 1,
  "depth": 10
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=market_liquidity&product_id={product_id}&depth={depth}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "market_liquidity",
  "product_id": 1,
  "depth": 10
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of spot / perp product for which to retrieve market liquidity.

depth

number

Yes

Number of price levels to retrieve. (`max: 100`)

## Response

**Note:**

* Each entry inside bids and asks is an array of price and size respectively. **Note**: that price is represented using fixed point, so it is `1e18` times greater than the decimal price.
* `timestamp` is in nanoseconds.


Last updated 1 month ago

---


# Market Prices

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/market-prices

## Rate limits

* 2400 requests/min or 40 requests/sec per IP address. (**weight = 1**) or length of `product_ids` for [multi-product markets](/developer-resources/api/gateway/queries/market-prices#multiple-products) query.

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Single Product

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "market_price",
  "product_id": 1
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=market_price&product_id={product_id}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "market_price",
  "product_id": 1
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of spot / perp product for which to retrieve market price data.

## Response

**Note**: that price is represented using fixed point, so it is `1e18` times greater than the decimal price.

## Multiple Products

## Request

Websocket

REST

**Connect**

`WEBSOCKET [CORE_WEBSOCKET_ENDPOINT]`

**Message**

`POST /query`

**Body**

## Request Parameters

Parameter

Type

Required

Description

product\_ids

number[]

Yes

List of spot / perp products for which to retrieve market price data.

## Response


Last updated 1 month ago

---


# Max NLP Burnable

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/max-nlp-burnable

## Rate limits

* 120 requests/min or 20 requests every 10 seconds per IP address. (**weight = 20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "max_nlp_burnable",
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=max_nlp_burnable&sender={sender}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "max_nlp_burnable",
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000"
}
```

## Request Parameters

Parameter

Type

Required

Description

sender

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier.

## Response


Last updated 1 month ago

---


# Max NLP Mintable

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/max-nlp-mintable

## Rate limits

* 120 requests/min or 20 requests every 10 seconds per IP address. (**weight = 20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "max_nlp_mintable",
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000",
  "spot_leverage": "true"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=max_nlp_mintable&sender={sender}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "max_nlp_mintable",
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000",
  "spot_leverage": "true"
}
```

## Request Parameters

Parameter

Type

Required

Description

sender

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier.

spot\_leverage

boolean

No

Boolean sent as a string. indicates whether leverage should be used; when set to `false` , returns the max amount of base LP mintable possible without borrow. Defaults to `true`

## Response


Last updated 1 month ago

---


# Max Order Size

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/max-order-size

## Rate limits

* 480 requests/min or 80 requests every 10 seconds per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "max_order_size",
  "product_id": 1,
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000",
  "price_x18": "23000000000000000000000",
  "direction": "short",
  "spot_leverage": "true",
  "reduce_only": "false",
  "isolated": "false"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=max_order_size&product_id={product_id}&sender={sender}&price_x18={price_x18}&direction={direction}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "max_order_size",
  "product_id": 1,
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000",
  "price_x18": "23000000000000000000000",
  "direction": "short",
  "spot_leverage": "true",
  "reduce_only": "false",
  "isolated": "false"
}
```

## Request Parameters

Parameter

Type

Required

Description

sender

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier.

product\_id

number

Yes

Id of spot / perp product for which to retrieve max order size.

price\_x18

string

Yes

An `int128` representing the price of the order multiplied by 1e18, sent as a string. For example, a price of 1 USDT0 would be sent as `"1000000000000000000"`

direction

string

Yes

`long` for max bid or `short` for max ask.

spot\_leverage

string

No

Boolean sent as a string. Indicates whether leverage should be used; when set to `false` , returns the max order possible without borrow. Defaults to `true`

reduce\_only

string

No

Boolean sent as a string. Indicates wether to retrieve the max order size to close / reduce a position. Defaults to `false`

isolated

string

No

Boolean sent as a string. When set to `true`, calculates max order size for an isolated margin position. Defaults to `false`. See [Isolated Margin](https://github.com/nadohq/nado-docs/blob/main/docs/basics/isolated-margin.md) to learn more.

## Response


Last updated 1 month ago

---


# Max Withdrawable

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/max-withdrawable

## Rate limits

* 480 requests/min or 80 requests every 10 seconds per IP address. (**weight = 5**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "max_withdrawable",
  "product_id": 1,
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000",
  "spot_leverage": "true"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=max_withdrawable&product_id={product_id}&sender={sender}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
  "type": "max_withdrawable",
  "product_id": 1,
  "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000",
  "spot_leverage": "true"
}
```

## Request Parameters

Parameter

Type

Required

Description

sender

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier.

product\_id

number

Yes

Id of spot / perp product for which to retrieve max withdrawable amount.

spot\_leverage

string

No

Boolean sent as a string. Indicates whether leverage should be used; when set to `false` , returns the max withdrawable amount possible without borrow. Defaults to `true`

## Response


Last updated 1 month ago

---


# NLP Locked Balances

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/nlp-locked-balances

## Rate limits

* 120 requests/min or 20 requests every 10 seconds per IP address. (**weight = 20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "nlp_locked_balances",
  "subaccount": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=nlp_locked_balances&subaccount={subaccount}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "nlp_locked_balances",
  "subaccount": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000"
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccount

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier.

## Response

## Response Fields

## NLP Locked Balances Response

Field name

Description

balance\_locked

Total balance that is currently locked (SpotBalance object)

balance\_unlocked

Total balance that is currently unlocked and available (SpotBalance object)

locked\_balances

Array of individual locked balance entries with their unlock times

## Locked Balance Entry

Field name

Description

balance

SpotBalance object containing the locked amount

unlocked\_at

Unix epoch timestamp (in seconds) when this balance will unlock

## SpotBalance Object

Field name

Description

product\_id

The product ID (typically 0 for USDT0/quote asset)

balance

Balance details object

balance.amount

The balance amount in x18 format (string)

balance.last\_cumulative\_funding\_x18

Last cumulative funding value in x18 format (string)

## Notes

* NLP positions have a 4-day lock period after minting before they can be burned (withdrawn)
* The `locked_balances` array shows individual lock entries, each with their own unlock timestamp
* `balance_locked` is the sum of all locked balances
* `balance_unlocked` represents balances that have passed their lock period and can be withdrawn


Last updated 1 month ago

---


# NLP Pool Info

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/nlp-pool-info

## Rate limits

* 120 requests/min or 20 requests every 10 seconds per IP address. (**weight = 20**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "nlp_pool_info"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=nlp_pool_info`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "nlp_pool_info"
}
```

## Request Parameters

This query does not require any parameters.

## Response

Copy

```
{
  "status": "success",
  "data": {
    "nlp_pools": [
      {
        "pool_id": 1,
        "subaccount": "0x0000000000000000000000000000000000000000000000000000000000000002",
        "owner": "0x1234567890123456789012345678901234567890",
        "balance_weight_x18": "500000000000000000",
        "subaccount_info": {
          "subaccount": "0x0000000000000000000000000000000000000000000000000000000000000002",
          "exists": true,
          "health": {
            "assets": "1000000000000000000000",
            "liabilities": "500000000000000000000",
            "initial_health": "250000000000000000000",
            "maintenance_health": "100000000000000000000"
          },
          "spot_balances": [],
          "perp_balances": []
        },
        "open_orders": []
      }
    ]
  },
  "request_type": "query_nlp_pool_info"
}
```

## Response Fields

## NLP Pool Info

Field name

Description

nlp\_pools

Array of NLP pool objects

## NLP Pool Object

Field name

Description

pool\_id

Unique identifier for the pool

subaccount

The subaccount address associated with this pool (bytes32 hex string)

owner

The owner address of the pool (bytes20 hex string)

balance\_weight\_x18

Weight of this pool's balance in x18 format (string representation of u128)

subaccount\_info

Complete subaccount information including health, balances, and positions

open\_orders

Array of currently open orders for this pool


Last updated 1 month ago

---


# Nonces

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/nonces

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "nonces",
  "address": "0x0000000000000000000000000000000000000000"
}
```

**GGET** `[GATEWAY_REST_ENDPOINT]/query?type=nonces&address={address}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "nonces",
  "address": "0x0000000000000000000000000000000000000000"
}
```

## Request Parameters

Parameter

Type

Required

Description

address

string

Yes

A `bytes20` sent as a hex string representing the wallet address.

## Response

**Note**: when doing any execute that is not `place_orders`, i.e. `withdraw_collateral`, `liquidate_subaccount`, you want to use `tx_nonce` as the nonce. `tx_nonce` increments by one each time a successful execute goes through. `order_nonce` is a historical artifact for the frontend, and simply returns the current timestamp in milliseconds plus 100000 multiplied by 2\*\*20.


Last updated 1 month ago

---


# Order

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/order

## Rate limits

* 2400 requests/min or 40 requests/sec per IP address. (**weight = 1**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "order",
  "product_id": 1,
  "digest": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=order&product_id={product_id}&digest={digest}`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "order",
  "product_id": 1,
  "digest": "0x0000000000000000000000000000000000000000000000000000000000000000"
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_id

number

Yes

Id of spot / perp product for which to retrieve order.

digest

string

Yes

Order digest to retrieve.

## Response

**Note**: that side of the order (buy/sell) is included in the sign of `amount` and `unfilled_amount` . They are positive if the order is a buy order, otherwise negative.


Last updated 1 month ago

---


# Status

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/status

## Rate limits

* 2400 requests/min or 40 requests/sec per IP address. (**weight = 1**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "status"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=status`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "status"
}
```

## Response

Copy

```
{
  "status": "success",
  "data": "active",
  "request_type": "query_status",
}
```

The offchain sequencer could be in any of the following statuses:

* `active`: accepting incoming executes.
* `failed`: sequencer is in a failed state.


Last updated 1 month ago

---


# Subaccount Info

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/subaccount-info

## Rate limits

The rate limit weight varies based on the request parameters:

* **Basic query** (no `txns`): **weight = 2**

  + 1200 requests/min or 200 requests every 10 seconds per IP address
* **With simulation** (`txns` provided): **weight = 10**

  + 240 requests/min or 40 requests every 10 seconds per IP address
* **With simulation + pre\_state** (`txns` and `pre_state="true"`): **weight = 15**

  + 160 requests/min or ~26 requests every 10 seconds per IP address

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
    "type": "subaccount_info",
    "subaccount": "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000",
    "txns": "[{\"apply_delta\":{\"product_id\":4,\"subaccount\":\"0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000\",\"amount_delta\":\"10790000000000000000\",\"v_quote_delta\":\"-35380410000000000000000\"}}]"
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=subaccount_info&subaccount={subaccount}&txns=[{"apply_delta":{"product_id":2,"subaccount":"0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000","amount_delta":"100000000000000000","v_quote_delta":"3033500000000000000000"}}]`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
    "type": "subaccount_info",
    "subaccount": "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000",
    "txns": "[{\"apply_delta\":{\"product_id\":4,\"subaccount\":\"0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000\",\"amount_delta\":\"10790000000000000000\",\"v_quote_delta\":\"-35380410000000000000000\"}}]"
}
```

## Request Parameters

Parameter

Type

Required

Description

subaccount

string

Yes

A `bytes32` sent as a hex string; includes the address and the subaccount identifier. See [sender field structure](/developer-resources/api/gateway/executes#sender-field-structure) for details.

txns

string

no

A list of transactions to get an estimated/simulated view. see more info below.

pre\_state

string

no

When `"true"` and `txns` are provided, returns the subaccount state before the transactions were applied in the `pre_state` field. Defaults to `"false"`.

## Supported txs for an estimated subaccount info

The following are the supported `txns` you can provide to get an estimated view of your subaccount.

**Note**: these `txns` are only used to simulate what your subaccount would look like if they were executed.

## ApplyDelta

Updates internal balances for the `product_id` and amount deltas provided.

## Response

**Note**:

* `healths`:

  + `healths[0]`: info about your initial health, which is weighted by `long_weight_initial_x18` and `short_weight_initial_x18.`
  + `healths[1]`: info about your maintenance health, which is weighted by `long_weight_maintenance_x18` and `short_weight_maintenance_x18.`
  + `healths[2]`: info about your unweighted health.
* `health_contributions` is indexed by product\_id and represents the contribution of the corresponding product to the final health.

  + `health_contributions[product_id][0]``: contribution to healths[0]`
  + `health_contributions[product_id][1]``: contribution to healths[1]`
  + `health_contributions[product_id][2]``: contribution to healths[2]`
* `pre_state`: (Optional) When `pre_state="true"` is provided with `txns`, this field contains the subaccount state **before** the simulated transactions were applied. This allows you to compare the before/after states when simulating transactions.

  + `pre_state.healths`: Same structure as the main `healths` field, but reflecting the state before transactions
  + `pre_state.health_contributions`: Health contributions before transactions
  + `pre_state.spot_balances`: Spot balances before transactions
  + `pre_state.perp_balances`: Perpetual balances before transactions

## Example with `pre_state`

When you want to simulate transactions and compare the before/after states, you can use the `pre_state` parameter:

## Request

REST (GET)

REST (POST)

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=subaccount_info&subaccount={subaccount}&txns=[{"apply_delta":{"product_id":2,"subaccount":"0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000","amount_delta":"100000000000000000","v_quote_delta":"3033500000000000000000"}}]&pre_state="true"`

`POST [GATEWAY_REST_ENDPOINT]/query`

## Response

The response will now include a `pre_state` field showing the state before the simulated transactions:

**Use Case**: The `pre_state` feature is particularly useful for:

* **Position Simulation**: Preview how a potential trade would affect your health and balances
* **Risk Analysis**: Compare health metrics before and after simulated transactions
* **UI/UX**: Display "before → after" views to users when they're about to execute trades
* **Testing**: Validate transaction impacts without executing them on-chain


Last updated 1 month ago

---


# Symbols

Source: https://docs.nado.xyz/developer-resources/api/gateway/queries/symbols

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Websocket

REST (GET)

REST (POST)

**Connect**

`WEBSOCKET [GATEWAY_WEBSOCKET_ENDPOINT]`

**Message**

Copy

```
{
  "type": "symbols",
  "product_ids": [1, 2]
}
```

**GET** `[GATEWAY_REST_ENDPOINT]/query?type=symbols&product_type=spot`

`POST [GATEWAY_REST_ENDPOINT]/query`

**Message**

Copy

```
{
  "type": "symbols",
  "product_ids": [1, 2, 3, 4],
  "product_type": "spot"
}
```

## Request Parameters

Parameter

Type

Required

Description

product\_ids

number[]

No

An array of product ids. Only available for POST and WS requests.

product\_type

string

No

Type of products to return, must be:
"spot" | "perp".

## Response

**Note**:

* All products have are quoted against USDT0, except for product 0.

## Response fields

## Symbols

All numerical values are returned as strings and scaled by 1e18.

Field name

Description

type

Product type, "spot" or "perp"

product\_id

Product id

symbol

Product symbol

price\_increment\_x18

Price increment, a.k.a tick size

size\_increment

Size increment, in base units

min\_size

Minimum order size, in base units

maker\_fee\_rate\_x18

Maker fee rate, given as decimal rate

taker\_fee\_rate\_x18

Taker fee rate, given as decimal rate

long\_weight\_initial\_x18

Long initial margin weight, given as decimal

long\_weight\_maintenance\_x18

Long maintenance margin weight, given as decimal

max\_open\_interest\_x18

Maximum open interest, null if no limit


Last updated 1 month ago

---


# Signing

Source: https://docs.nado.xyz/developer-resources/api/gateway/signing

All executes are signed using [EIP712](https://eips.ethereum.org/EIPS/eip-712). Each execute request contains:

1. A piece of structured data that includes the sender address i.e: the `primaryType` that needs to be signed.
2. A signature of the hash of that structured data, signed by the sender.

## Domain

The following is the domain required as part of the EIP712 structure:

Copy

```
{
    name: 'Nado',
    version: '0.0.1',
    chainId: chainId,
    verifyingContract: contractAddress
}
```

You can retrieve the corresponding chain id and verifying contract via the [contracts](/developer-resources/api/gateway/queries/contracts) query.

**Note**: make sure to use the correct verifying contract for each execute:

* For place order: should use `address(producId)` i.e: the 20 bytes hex representation of the `productId` for the order. For example, the verify contract of product `18` is `0x0000000000000000000000000000000000000012` .
* For everything else: should use the endpoint address.

See more details in the [contracts](/developer-resources/api/gateway/queries/contracts) query page.

Copy

```
def gen_order_verifying_contract(product_id: int) -> str:
    """
    Generates the order verifying contract address based on the product ID.

    Args:
        product_id (int): The product ID for which to generate the verifying contract address.

    Returns:
        str: The generated order verifying contract address in hexadecimal format.
    """
    be_bytes = product_id.to_bytes(20, byteorder="big", signed=False)
    return "0x" + be_bytes.hex()
```

## EIP712 Types

See below the EIP712 type for each execute:

See more details in the **Signing** section of each execute's page.

## [Place Order](/developer-resources/api/gateway/executes/place-order)

**Primary Type**: `Order`

Solidity struct that needs to be signed:

**JSON representation:**

## [Cancel Orders](/developer-resources/api/gateway/executes/cancel-orders)

**Primary Type:** `Cancellation`

Solidity struct that needs to be signed:

**JSON representation:**

## [Cancel Product Orders](/developer-resources/api/gateway/executes/cancel-product-orders)

**Primary Type**: `CancellationProducts`

Solidity struct that needs to be signed:

**JSON representation:**

## [Withdraw Collateral](/developer-resources/api/gateway/executes/withdraw-collateral)

**Primary Type:** `WithdrawCollateral`

Solidity struct that needs to be signed:

**JSON representation:**

## [Liquidate Subaccount](/developer-resources/api/gateway/executes/liquidate-subaccount)

**Primary Type:** `LiquidateSubaccount`

Solidity struct that needs to be signed:

**JSON representation:**

## [Mint NLP](/developer-resources/api/gateway/executes/mint-nlp)

**Primary Type**: `MintNlp`

Solidity struct that needs to be signed:

**JSON representation:**

## [Burn NLP](/developer-resources/api/gateway/executes/burn-nlp)

**Primary Type:** `BurnNlp`

Solidity struct that needs to be signed:

**JSON representation:**

## [Link Signer](/developer-resources/api/gateway/executes/link-signer)

**Primary Type**: `LinkSigner`

Solidity struct that needs to be signed:

**JSON representation:**

## [List Trigger Orders](/developer-resources/api/trigger/queries/list-trigger-orders)

**Primary Type**: `ListTriggerOrders`

Solidity struct that needs to be signed:

**JSON representation:**

## [Authenticate Subscription Streams](/developer-resources/api/subscriptions#authentication)

**Primary Type**: `StreamAuthentication`

Struct that needs to be signed:

**JSON representation:**


Last updated 11 days ago

---


# Examples

Source: https://docs.nado.xyz/developer-resources/api/gateway/signing/examples

The following are full examples of EIP12 typed data for each of Nado's executes. Each execute includes a `sender` field which is a solidity `bytes32` . There are two components to this field:

* an `address` that is a `bytes20`
* a subaccount identifier that is a `bytes12`

For example, if your address was `0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43`, and you wanted to use the default subaccount identifier (i.e: an empty identifier `""`) you can set `sender` to `0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43000000000000000000000000` , which sets all bytes of the subaccount identifier to `0`.

**Note**: a `bytes32` representation of the sender must used when signing the request.

See below a sample util to convert a hex to a **bytes32**:

Python

Typescript

Copy

```
def hex_to_bytes32(hex_string):
    if hex_string.startswith("0x"):
        hex_string = hex_string[2:]
    data_bytes = bytes.fromhex(hex_string)
    padded_data = data_bytes + b"\x00" * (32 - len(data_bytes))
    return padded_data

sender = hex_to_bytes32('0x841fe4876763357975d60da128d8a54bb045d76a64656661756c740000000000')
```

Copy

```
import { arrayify } from 'ethers/lib/utils';

export function hexToBytes32(subaccount: string) {
  const subaccountBytes = arrayify(subaccount);
  const bytes32 = new Uint8Array(32);
  for (let i = 0; i < Math.min(subaccountBytes.length, 32); i++) {
    bytes32[i] = subaccountBytes[i];
  }
  return bytes32;
}

const sender = hexToBytes32('0x841fe4876763357975d60da128d8a54bb045d76a64656661756c740000000000')
```

## EIP712 Typed data examples

Place Order

Cancel Orders

Cancel Product orders

Link Signer

Copy

```
{
    'types': {
        'EIP712Domain': [
            {'name': 'name', 'type': 'string'},
            {'name': 'version', 'type': 'string'},
            {'name': 'chainId', 'type': 'uint256'},
            {'name': 'verifyingContract', 'type': 'address'}
        ],
        'Order': [
            {'name': 'sender', 'type': 'bytes32'},
            {'name': 'priceX18', 'type': 'int128'},
            {'name': 'amount', 'type': 'int128'},
            {'name': 'expiration', 'type': 'uint64'},
            {'name': 'nonce', 'type': 'uint64'},
            {'name': 'appendix', 'type': 'uint128'},
        ],
    },
    'primaryType': 'Order',
    'domain': {
        'name': 'Nado',
        'version': '0.0.1',
        'chainId': 763373,  
        'verifyingContract': '0x0000000000000000000000000000000000000001'
    },
    'message': {
        'sender': hex_to_bytes32('0x841fe4876763357975d60da128d8a54bb045d76a64656661756c740000000000'),
        'priceX18': 28898000000000000000000,
        'amount': -10000000000000000,
        'expiration': 4611687701117784255,
        'appendix': 1537,  # Version 1, POST_ONLY order
        'nonce': 1764428860167815857,
    },
}
```

Withdraw Collateral

Liquidate Subaccount

Mint NLP

Burn NLP

List Trigger Orders


Last updated 12 days ago

---


# Q&A

Source: https://docs.nado.xyz/developer-resources/api/gateway/signing/q-and-a

## Q: **What is Nado's EIP712 domain?**

Copy

```
{
    name: 'Nado',
    version: '0.0.1',
    chainId: chainId,
    verifyingContract: contractAddress
}
```

See [signing](/developer-resources/api/gateway/signing#domain) for more details.

## Q: How can i retrieve the verifying contracts to use?

* Via the [contracts](/developer-resources/api/gateway/queries/contracts) query for all executes except place orders.

## Q: Which contract should I use for each execute?

* For place orders: must be computed as `address(productId)`. For example, the verify contract of product `18` is `0x0000000000000000000000000000000000000012`.
* For everything else: use the endpoint contract from the contracts query.

See the [contracts](/developer-resources/api/gateway/queries/contracts#response) query for more details.

## Q: I am running into signature errors, how to fix?

Signature errors can arise for several reasons:

* **An invalid struct**: confirm you are signing the correct struct. See the [Signing](/developer-resources/api/gateway/signing) page to verify the struct of each execute request.
* **An invalid chain id**: confirm you have the correct chain id for the network you are on.
* **An invalid verifying contract**: confirm you have the correct verifying contract address for the network and execute you are signing. i.e: confirm you are using the correct orderbook address for place orders and endpoint address for everything else.

## Q: Is any other signing standard supported?

No, only [EIP712](https://eips.ethereum.org/EIPS/eip-712).

## Q: Are there any examples you can provide?

See [examples](/developer-resources/api/gateway/signing/examples).

## Q: What is the PrimaryType of execute X?

All primary types are listed in our [signing](/developer-resources/api/gateway/signing) page.


Last updated 1 month ago

---


# Integrate via Smart Contracts

Source: https://docs.nado.xyz/developer-resources/api/integrate-via-smart-contracts

Smart contracts can use the `LinkSigner` transaction type (see [Link Signer](/developer-resources/api/gateway/executes/link-signer)) to perform the following:

1. Deposit into Nado.
2. LinkSigner an externally owned account ([EOA](https://ethereum.org/en/developers/docs/accounts/#externally-owned-accounts-and-key-pairs)).
3. Have the externally owned account trade using the smart contract's assets on Nado.

## Setup: Depositing into Nado + Linking an EOA

1. Deposits are always on-chain, as such, users can simply have their smart contract call `depositCollateral` on our `Endpoint` contract (see [Contracts](https://github.com/nadohq/nado-docs/blob/main/docs/developer-resources/contracts.md) for addresses).
2. The contract needs to have 1 USDT0 available to pay for slow-mode fee and approve the endpoint contract, assemble the bytes for a slow mode linked signer transaction, and submit it via [submitSlowModeTransaction](https://github.com/nadohq/nado-contracts/blob/main/core/contracts/Endpoint.sol).

You can find the requisite parsing logic in the [Endpoint](https://github.com/nadohq/nado-contracts/blob/main/core/contracts/Endpoint.sol) contract.

## Example

Copy

```
struct LinkSigner {
    bytes32 sender;
    bytes32 signer;
    uint64 nonce;
}

function linkNadoSigner(
        address nadoEndpoint,
        address externalAccount,
        address usdt0Address
    ) external {
    // 1. a slow mode fee of 1 USDT0 needs to be avaliable and approved
    ERC20 usdt0Token =  ERC20(usdt0Address);
    
    // NOTE: should double check the USDT0 decimals in the corresponding chain.
    // e.g: it's 1e6 on arbitrum, whereas it's 1e18 on blast, etc.
    uint256 SLOW_MODE_FEE = 1e6;
    usdt0Token.transferFrom(msg.sender, address(this), SLOW_MODE_FEE);
    usdt0Token.approve(nadoEndpoint, SLOW_MODE_FEE);
    
    // 2. assamble the link signer slow mode transaction
    bytes12 defaultSubaccountName = bytes12(abi.encodePacked("default"));
    bytes32 contractSubaccount = bytes32(
        abi.encodePacked(uint160(address(this)), defaultSubaccountName)
    );
    bytes32 externalSubaccount = bytes32(
        uint256(uint160(externalAccount)) << 96
    );
    LinkSigner memory linkSigner = LinkSigner(
        contractSubaccount,
        externalSubaccount,
        IEndpoint(nadoEndpoint).getNonce(contractSubaccount)
    );
    bytes memory txs = abi.encodePacked(
        uint8(13),
        abi.encode(linkSigner)
    );
    
    // 3. submit slow mode transaction
    IEndpoint(nadoEndpoint).submitSlowModeTransaction(txs);
}
```

Once the transaction is confirmed, it may take a few seconds for it to make its way into the Nado offchain sequencer. Afterwards, you can sign transactions that have sender `contractSubaccount` using `externalSubaccount`, and they will be accepted by the sequencer and the blockchain.


Last updated 17 days ago

---


# Order Appendix

Source: https://docs.nado.xyz/developer-resources/api/order-appendix

The **Order Appendix** is a 128-bit integer that encodes extra order parameters like execution type, isolated margin, and trigger configurations.

## Bit Layout

Copy

```
| value   | reserved | trigger | reduce only | order type | isolated | version |
| 64 bits | 50 bits  | 2 bits  | 1 bit       | 2 bits     | 1 bit    | 8 bits  |
| 127..64 | 63..14   | 13..12  | 11          | 10..9      | 8        | 7..0    |
```

## Fields (from LSB to MSB)

## Version

**8-bits (0-7)**. Protocol version identifier. Currently `1`. May increment when encoding structure updates.

## Isolated

**1-bit (8)**. Indicates whether the order uses isolated margin. Isolated positions have dedicated margin for a specific product, creating a separate isolated subaccount. The original account becomes the "parent subaccount" that can manage the isolated position.

*Key Properties:*

* Creates isolated subaccount with dedicated margin
* Only quote transfers allowed between isolated and parent subaccounts
* Parent account can sign orders for isolated subaccount
* Cannot be combined with TWAP orders

*Example:*

## Order Type

**2-bits (9-10)**. Execution behavior for the order.

*Values:*

* `0` - `DEFAULT`: Standard limit order behavior.
* `1` - `IOC (Immediate or Cancel)`: Execute immediately, cancel unfilled portion.
* `2` - `FOK (Fill or Kill)`: Execute completely or cancel entire order.
* `3` - `POST_ONLY`: Only add liquidity, reject if would take liquidity.

*Example:*

## Reduce Only

**1-bit (11)**. Restricts order to only decrease existing positions. Prevent accidentally increasing position size. Order will be rejected if it would increase the position in the same direction.

*Use Cases:*

* Risk management when closing positions.
* Taking profits without adding exposure.
* Automated position reduction strategies.

*Example:*

## Trigger Type

**2-bits (12-13)**. Conditional execution behavior.

*Values:*

* `0` - `NONE`: Execute immediately (regular order).
* `1` - `PRICE`: Price-based conditional order.
* `2` - `TWAP`: Time-Weighted Average Price execution.
* `3` - `TWAP_CUSTOM_AMOUNTS`: TWAP with randomized amounts.

*Example:*

## Reserved

**50-bits (14-63)**. Reserved for future protocol extensions. Must be set to `0`.

## Value

**64-bits (64-127)**. Context-dependent data based on other flags.

## **TWAP Configuration (when trigger = 2 or 3)**

Encodes TWAP execution parameters in the 64-bit value field:

**Fields:**

* `times`: Number of TWAP executions.
* `slippage_x6`: Maximum slippage × 1\_000\_000 (6 decimal precision).

**Example:**

## **Isolated Margin (when isolated = 1)**

Amount of quote (margin\_x6) to transfer to isolated subaccount on first fill, stored in the 64-bit value field.

**Important:** Isolated margin is stored in **x6 precision** (6 decimals) in the appendix value field.

* Stored as `margin_x6` (6 decimal places)
* Takes up 64 bits (bits 64-127 of the appendix)

*Example:*

## Constraints

* **Isolated + TWAP**: Cannot combine isolated orders with TWAP (trigger types 2 or 3).
* **TWAP Requirements**: TWAP orders must specify both `twap_times` and `twap_slippage_frac` .
* **Isolated Margin**: Can only set `isolated_margin` when `isolated=True` .

## Migration from Legacy Format

**Before (deprecated):**

* Order type encoded in `expiration` field.
* Reduce-only flag encoded in `nonce` field.
* Limited trigger functionality.

**After (current):**

* All flags consolidated in 128-bit `appendix` .
* `expiration` is pure timestamp.
* `nonce` encodes `recv_time` only.
* Enhanced trigger and isolated margin support.

## Building Appendix Values

## Using Python SDK (Recommended)

## Manual Bit Manipulation (Advanced)

Refer to [nado\_protocol.utils.order](https://nadohq.github.io/nado-python-sdk/_modules/nado_protocol/utils/order.html) for a detailed implementation.

## Utility Functions


Last updated 1 month ago

---


# Rate limits

Source: https://docs.nado.xyz/developer-resources/api/rate-limits

## Overview

* Nado uses a weight-based rate-limiting system across queries and executes. We limit based on `IP address`, `Wallet address`, and a global `max # of orders per subaccount per market`.
* These limits equally apply to both `http` requests and `Websocket` messages.
* Limits are applied on a `1 minute` and `10 seconds` basis.

## Limits

* IP addresses have a max weight limit of `2400` per minute or `400` every 10 seconds applied only to queries.
* Wallet addresses have a max weight limit of `600` per minute or `100` every 10 seconds applied only to executes.
* Users can have up to `500` open orders per subaccount per market.
* Orders have the following additional limits:

  + Place orders (with spot leverage): up to `600` per minute or `100` every 10 seconds across all markets.
  + Place orders (without spot leverage): up to `30` per minute or `5` every 10 seconds across all markets. **Note**: orders without spot leverage are `20x` more expensive to place due to additional health checks needed.
  + Order cancellations: up to `600` per minute or `100` every 10 seconds.

## Query Weights

Queries are rate-limited based on IP. The following weights are applied per query:

* [**Status**](/developer-resources/api/gateway/queries/status): `IP weight = 1`
* [**Contracts**](/developer-resources/api/gateway/queries/contracts): `IP weight = 1`
* [**Nonces**](/developer-resources/api/gateway/queries/nonces): `IP weight = 2`
* [**Order**](/developer-resources/api/gateway/queries/order)**:** `IP weight = 1`
* [**Orders**](/developer-resources/api/gateway/queries/orders): `IP weight = 2 * product_ids.length`
* [**Subaccount Info**](/developer-resources/api/gateway/queries/subaccount-info): `IP weight = 2` (or `10` with `txns`, or `15` with `txns` + `pre_state="true"`)
* [**Isolated Positions**](/developer-resources/api/gateway/queries/isolated-positions): `IP weight = 10`
* [**Market Liquidity**](/developer-resources/api/gateway/queries/market-liquidity): `IP weight = 1`
* [**Symbols**](/developer-resources/api/gateway/queries/symbols): `IP weight = 2`
* [**All Products**](/developer-resources/api/gateway/queries/all-products): `IP weight = 5`
* [**Edge All Products**](/developer-resources/api/gateway/queries/edge-all-products): `IP weight = 5`
* [**Market Prices**](/developer-resources/api/gateway/queries/market-prices)**:** `IP weight = product_ids.length`
* [**Max Order Size**](/developer-resources/api/gateway/queries/max-order-size)**:** `IP weight = 5`
* [**Max Withdrawable**](/developer-resources/api/gateway/queries/max-withdrawable)**:** `IP weight = 5`
* [**Max NLP Mintable**](/developer-resources/api/gateway/queries/max-nlp-mintable)**:** `IP weight = 20`
* [**Max NLP Burnable**](/developer-resources/api/gateway/queries/max-nlp-burnable)**:** `IP weight = 20`
* [**NLP Pool Info**](/developer-resources/api/gateway/queries/nlp-pool-info)**:** `IP weight = 20`
* [**NLP Locked Balances**](/developer-resources/api/gateway/queries/nlp-locked-balances)**:** `IP weight = 20`
* [**Health Groups**](/developer-resources/api/gateway/queries/health-groups)**:** `IP weight = 2`
* [**Linked Signer**](/developer-resources/api/gateway/queries/linked-signer)**:** `IP weight = 5`
* [**Insurance**](/developer-resources/api/gateway/queries/insurance): `IP weight = 2`
* [**Fee Rates**](/developer-resources/api/gateway/queries/fee-rates)**:** `IP weight = 2`
* [**Assets**](/developer-resources/api/v2/assets): `IP weight = 2`
* [**Orderbook**](/developer-resources/api/v2/orderbook): `IP weight = 1`

## Archive (indexer) Weights

* Archive (indexer) queries are rate-limited based on IP.
* IP addresses have a max weight limit of `2400` per minute or `400` every 10 seconds.

The following weights are applied per query:

* [**Orders**](/developer-resources/api/archive-indexer/orders)**:** `IP Weight = 2 + (limit * subaccounts.length / 20)`; where `limit` and `subaccounts` are query params.
* [**Matches**](/developer-resources/api/archive-indexer/matches)**:** `IP Weight = 2 + (limit * subaccounts.length / 10)`; where `limit` and `subaccounts` are query params.
* [**Events**](/developer-resources/api/archive-indexer/events)**:** `IP Weight = 2 + (limit * subaccounts.length / 10)`; where `limit` and `subaccounts` are query params.
* [**Candlesticks**](/developer-resources/api/archive-indexer/candlesticks)**:** `IP Weight = 1 + limit / 20`; where `limit` is a query param.
* [**Edge Candlesticks**](/developer-resources/api/archive-indexer/edge-candlesticks): `IP Weight = 1 + limit / 20`; where `limit` is a query param.
* [**Product Snapshots**](/developer-resources/api/archive-indexer/product-snapshots)**:** `IP Weight = 10` for single `products` query, or `10 * timestamps.length` for multiple `product_snapshots` query with max\_time parameter
* [**Funding Rate**](/developer-resources/api/archive-indexer/funding-rate)**:** `IP Weight = 2`
* [**Interest & funding payments**](/developer-resources/api/archive-indexer/interest-and-funding-payments)**:** `IP Weight = 5`
* [**Oracle Price**](/developer-resources/api/archive-indexer/oracle-price)**:** `IP Weight = 2`
* [**Oracle Snapshots**](/developer-resources/api/archive-indexer/oracle-snapshots): `IP Weight = max((snapshot_count * product_ids.length / 100), 2)`; where snapshot\_count is `interval.count.min(500)`
* [**Perp Prices**](/developer-resources/api/archive-indexer/perp-prices)**:** `IP Weight = 2` (includes both single `price` and multiple `perp_prices` queries)
* [**Market Snapshots**](/developer-resources/api/archive-indexer/market-snapshots)**:** `IP Weight = max((snapshot_count * product_ids.length / 100), 2)`; where snapshot\_count is `interval.count.min(500)`
* [**Edge Market Snapshots**](/developer-resources/api/archive-indexer/edge-market-snapshots): `IP weight = (interval.count.min(500) / 20) + (interval.count.clamp(2, 20) * 2)`
* [**Subaccounts**](/developer-resources/api/archive-indexer/subaccounts)**:** `IP Weight = 2`
* [**Subaccount Snapshots**](/developer-resources/api/archive-indexer/subaccount-snapshots): `IP Weight = 2 + (limit * subaccounts.length / 10)`; where `limit` and `subaccounts` are query params.
* [**Linked Signers**](/developer-resources/api/archive-indexer/linked-signers): `IP Weight = 2`
* [**Linked Signer Rate Limit**](/developer-resources/api/archive-indexer/linked-signer-rate-limit)**:** `IP Weight = 2`
* [**Isolated Subaccounts**](/developer-resources/api/archive-indexer/isolated-subaccounts): `IP Weight = 2`
* [**Signatures**](/developer-resources/api/archive-indexer/signatures): `IP Weight = 2 + len(digests) / 10`; where `digests` is a query param.
* [**Fast Withdrawal Signature**](/developer-resources/api/archive-indexer/fast-withdrawal-signature): `IP Weight = 10`
* [**NLP Funding Payments**](/developer-resources/api/archive-indexer/nlp-funding-payments): `IP Weight = 5`
* [**NLP Interest Payments**](/developer-resources/api/archive-indexer/nlp-interest-payments): `IP Weight = 5`
* [**NLP Snapshots**](/developer-resources/api/archive-indexer/nlp-snapshots): `IP Weight = limit.min(500) / 100`; where `limit` is a query param.
* [**Tx Hashes**](/developer-resources/api/archive-indexer/tx-hashes): `IP Weight = idxs.length * 2`; where `idxs` is an array of submission indices (max 100).
* [**Liquidation Feed**](/developer-resources/api/archive-indexer/liquidation-feed)**:** `IP Weight = 2`
* [**Sequencer Backlog**](/developer-resources/api/archive-indexer/sequencer-backlog): `IP Weight = 1`
* [**Direct Deposit Address**](/developer-resources/api/archive-indexer/direct-deposit-address): `IP Weight = 10`
* [**Quote Price**](/developer-resources/api/archive-indexer/quote-price): `IP Weight = 2`
* [**Ink Airdrop**](/developer-resources/api/archive-indexer/ink-airdrop): `IP Weight = 2`

## Execute Weights

Executes are rate-limited based on Wallet address. The following weights are applied per execute:

* [**Place order**](/developer-resources/api/gateway/executes/place-order)**:**

  + With spot leverage: `Wallet weight = 1`
  + Without spot leverage: `Wallet weight = 20`
* [**Place orders**](/developer-resources/api/gateway/executes/place-orders)**:**

  + With spot leverage: `Wallet weight = 1 per order`
  + Without spot leverage: `Wallet weight = 20 per order`
  + **Note**: 50ms processing penalty per request
* [**Cancel orders**](/developer-resources/api/gateway/executes/cancel-orders)**:**

  + When no **digests** are provided: `Wallet weight = 1`
  + When **digests** are provided: `Wallet weight = total digests`
* [**Cancel Product Orders**](/developer-resources/api/gateway/executes/cancel-product-orders)**:**

  + When no **productIds** are provided**:** `Wallet weight = 50`
  + When **productIds** are provided: `Wallet weight = 5 * total productIds`
* [**Cancel And Place**](/developer-resources/api/gateway/executes/cancel-and-place):

  + The sum of [Cancel orders](/developer-resources/api/gateway/executes/cancel-orders) + [Place order](/developer-resources/api/gateway/executes/place-order) limits
* [**Withdraw Collateral**](/developer-resources/api/gateway/executes/withdraw-collateral)**:**

  + With spot leverage: `Wallet weight = 10`
  + Without spot leverage: `Wallet weight = 20`
* [**Liquidate Subaccount**](/developer-resources/api/gateway/executes/liquidate-subaccount): `Wallet weight = 20`
* [**Mint NLP**](/developer-resources/api/gateway/executes/mint-nlp): `Wallet weight = 10`
* [**Burn NLP**](/developer-resources/api/gateway/executes/burn-nlp): `Wallet weight = 10`
* [**Link Signer**](/developer-resources/api/gateway/executes/link-signer): `Wallet weight = 30`

  + Can only perform a max of 50 link signer requests every 7 days per subaccount.
* [**Transfer Quote:**](/developer-resources/api/gateway/executes/transfer-quote) `Wallet weight = 10`

  + Can only transfer to a max of 5 new recipients within 24hrs.

## Trigger Service Limits

The trigger service has additional limits specific to conditional orders:

* **Pending trigger orders**: Max of `25` pending trigger orders per product per subaccount
* **TWAP orders**: Must use IOC execution type and cannot be combined with isolated margin


Last updated 26 days ago

---


# Subscriptions

Source: https://docs.nado.xyz/developer-resources/api/subscriptions

## Overview

To interact with the subscription API, send websocket messages to `WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`.

Subscription connections must set the `Sec-WebSocket-Extensions` header to include `permessage-deflate`.

## Endpoints

## Testnet:

* `wss://gateway.test.nado.xyz/v1/subscribe`

**Note**: You must send ping frames every 30 seconds to keep the websocket connection alive.

[Authentication](/developer-resources/api/subscriptions/authentication)[Streams](/developer-resources/api/subscriptions/streams)

Last updated 1 month ago

---


# Authentication

Source: https://docs.nado.xyz/developer-resources/api/subscriptions/authentication

## Rate limits

A **single wallet address** can be authenticated by up to 5 websocket connections, regardless of the originating IP address. Connections exceeding these limits will be automatically disconnected.

See [rate limits](/developer-resources/api/subscriptions/rate-limits) for more details.

## Request

To access streams that require authentication, submit a request with the `method` field set to `authenticate`.

Authenticate

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

Copy

```
{
  "method": "authenticate",
  "id": 0,
  "tx": {
    "sender": "0x...",
    "expiration": "1..."
  },
  "signature": "0x..."
}
```

## Request Parameters

Parameter

Type

Required

Description

method

string

Yes

`authenticate`

id

number

Yes

Can be set to any positive integer. Can be used to identify the websocket request / response.

tx

object

Yes

`StreamAuthentication` object that needs to be signed. See [Signing](/developer-resources/api/subscriptions/authentication#signing) section for more details.

tx.sender

string

Yes

A hex string representing a `bytes32` of a specific subaccount.

tx.expiration

string

Yes

Represents the expiration time in milliseconds since the Unix epoch.

signature

string

Yes

Hex string representing hash of the **signed** `StreamAuthentication` object.See [Signing](/developer-resources/api/subscriptions/authentication#signing) section for more details.

**Notes**:

* Although sender specifies a specific subaccount, authentication applies to the entire wallet address, enabling access to authenticated streams for different subaccounts under that address.
* Once authenticated, the authentication status of that websocket connection cannot be changed and stays for the duration of the connection.

## Signing

See more details and examples in our [signing](/developer-resources/api/gateway/signing) page.

The typed data struct that needs to be signed is:

`sender`: A hex string representing a `bytes32` of a specific subaccount. The signature must be signed by the wallet address specified by sender.

`expiration`: Represents the expiration time in milliseconds since the Unix epoch. Requests will be denied if the expiration is either smaller than the current time or more than 100 seconds ahead of it.

**Notes**:

* Should use the endpoint address as `verifyingContract`.
* For signing, you should always use the data type specified in the typed data struct which might be different from the type sent in the request e.g: `expiration` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## **Response**


Last updated 1 month ago

---


# Streams

Source: https://docs.nado.xyz/developer-resources/api/subscriptions/streams

## Available Streams

See below the available streams you can subscribe to:

Copy

```
pub enum StreamSubscription {
    // pass `null` product_id to subscribe to all products
    OrderUpdate { product_id: Option<u32>, subaccount: H256 },
    Trade { product_id: u32 },
    BestBidOffer { product_id: u32 },
    // pass `null` product_id to subscribe to all products
    Fill { product_id: Option<u32>, subaccount: H256 },
    // pass `null` product_id to subscribe to all products
    PositionChange { product_id: Option<u32>, subaccount: H256},
    BookDepth { product_id: u32 },
    // pass `null` product_id to subscribe to all products
    Liquidation { product_id: Option<u32> },
    LatestCandlestick {
        product_id: u32,
        // time interval in seconds (e.g., 60, 300, 900, 3600)
        granularity: i32
    },
    FundingPayment { product_id: u32 },
    // pass `null` product_id to subscribe to all products
    FundingRate { product_id: Option<u32> }
}
```

## **Subscribing to a stream**

Order Update

Trade

Best Bid Offer

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: Yes.*

**Note**: Set `product_id` to `null` to subscribe to order updates across all products for the subaccount.

**Subscribe to all products:**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

Fill

Position Change

Book Depth

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

**Note**: Set `product_id` to `null` to subscribe to fills across all products for the subaccount.

**Subscribe to all products:**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

**Note**: Set `product_id` to `null` to subscribe to position changes across all products for the subaccount.

**Subscribe to all products:**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

Liquidation

Latest Candlestick

Funding Payment

Funding Rate

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

**Note**: Set `product_id` to `null` to subscribe to liquidations across all products.

**Subscribe to all products:**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

See all supportes `granularity` values in [Available Granularities](/developer-resources/api/archive-indexer/candlesticks#available-granularities)

*Requires Authentication: No.*

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

*Requires Authentication: No.*

**Note**: Set `product_id` to `null` to subscribe to funding rate updates across all products.

**Subscribe to all products:**

## **Response**

## **Unsubscribing**

Order Update

Trade

Best Bid Offer

Fill

Position Change

Book Depth

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

**Unsubscribe from all products:**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

**Connect**

`WEBSOCKET [SUBSCRIPTIONS_ENDPOINT]`

**Message**

## **Response**

## **Listing subscribed streams**

## Response


Last updated 1 month ago

---


# Trigger

Source: https://docs.nado.xyz/developer-resources/api/trigger

The trigger service enables sophisticated order execution strategies through conditional triggers:

## Order Types

## **Price Triggers**

Execute orders when price conditions are met:

* **Stop orders**: Trigger when price moves above or below a threshold
* **Take profit/Stop loss**: Automated position management
* **Support multiple price sources**: Oracle price, last trade price, or mid-book price

## **Time Triggers (TWAP)**

Execute large orders over time using Time-Weighted Average Price:

* **Split large orders**: Break into smaller executions to reduce market impact
* **Configurable intervals**: Set time between executions
* **Slippage protection**: Built-in slippage limits for each execution
* **Custom amounts**: Specify exact amounts for each execution or split evenly

## API Structure

There are two types of actions:

* `Execute`: Modifies state (place/cancel orders)
* `Query`: Fetches information (list orders, TWAP status)

**HTTP Endpoints:**

* `POST [TRIGGER_ENDPOINT]/execute` for order placement and cancellation
* `POST [TRIGGER_ENDPOINT]/query` for querying trigger order status

`HTTP` requests must set the `Accept-Encoding` to include `gzip`, `br` or `deflate`

## Rate Limits

* **Maximum pending orders**: 25 pending trigger orders per product per subaccount
* **TWAP constraints**: Must use IOC execution type, cannot combine with isolated margin

## Key Requirements

## **Order Appendix Configuration**

All trigger orders require proper [Order Appendix](/developer-resources/api/order-appendix) configuration:

* **Trigger type**: Specify price (1), TWAP (2), or TWAP with custom amounts (3) in appendix bits
* **Execution type**: TWAP orders **must** use IOC execution
* **TWAP parameters**: Encode execution count and slippage limits in appendix value field

## Endpoints

## Testnet:

* `https://trigger.test.nado.xyz/v1`



Last updated 1 month ago

---


# List Trigger Orders

Source: https://docs.nado.xyz/developer-resources/api/trigger/queries/list-trigger-orders

## Request

Basic query

Fetch by digest

Filter by type and status

Filter by reduce-only

`POST [TRIGGER_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "list_trigger_orders",
  "tx": {
    "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
    "recvTime": "1688768157050"
  },
  "signature": "0x",
  "product_ids": [1, 2],
  "max_update_time": 1688768157,
  "limit": 20
}
```

`POST [TRIGGER_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "list_trigger_orders",
  "tx": {
    "sender": "0x7a5ec2748e9065794491a8d29dcf3f9edb8d7c43746573743000000000000000",
    "recvTime": "1688768157050"
  },
  "signature": "0x",
  "digests": ["0x5886d5eee7dc4879c7f8ed1222fdbbc0e3681a14c1e55d7859515898c7bd2038"],
  "limit": 20
}
```

`POST [TRIGGER_ENDPOINT]/query`

**Body**

`POST [TRIGGER_ENDPOINT]/query`

**Body**

## Request Parameters

**Note**: `max_update_time` It's the time that the trigger order last changed state. For example, if a trigger order is placed & pending, the update time = time of placement. If the trigger order is cancelled, then the update time = time of cancellation.

Parameter

Type

Required

Description

tx

object

Yes

List trigger orders transaction object. See [Signing](/developer-resources/api/trigger/queries/list-trigger-orders#signing) section for details on the transaction fields.

tx.sender

string

Yes

Hex string representing the subaccount's 32 bytes (address + subaccount name) of the tx sender.

tx.recvTime

string

Yes

Encoded time in milliseconds after which the list trigger orders transaction will be ignored. cannot be more than 100 seconds from the time it is received by the server.

signature

string

Yes

Signed transaction. See [Signing](/developer-resources/api/trigger/queries/list-trigger-orders#signing) section for more details.

product\_ids

number[]

No

If provided, returns trigger orders for the specified products; otherwise, returns trigger orders for all products.

trigger\_types

string[]

No

If provided, filters by trigger type. Values: `price_trigger`, `time_trigger`.

status\_types

string[]

No

If provided, filters by order status. Values: `cancelled`, `triggered`, `internal_error`, `triggering`, `waiting_price`, `waiting_dependency`, `twap_executing`, `twap_completed`.

max\_update\_time

number

No

If provided, returns all trigger orders that were last updated up to `max_update_time`. must be a unix epoch in seconds.

max\_digest

string

No

If provided, returns all trigger orders up to the given order digest (exclusive). This can be used for pagination.

digests

string[]

No

If provided, only returns the trigger orders for the associated digests. **Note**: all other filters are ignored when `digests` is provided.

reduce\_only

boolean

No

If provided, filters trigger orders by reduce-only flag. `true` returns only orders that can only decrease existing positions. If omitted, returns all orders regardless of reduce-only status.

limit

number

No

If provided, returns the most recently updated trigger orders up to `limit`. defaults to 100. max limit is 500.

## Signing

See more details and and examples in our [signing](/developer-resources/api/gateway/signing) page.

The solidity typed data struct that needs to be signed is:

`sender`: a `bytes32` sent as a hex string; includes the address and the subaccount identifier

`recvTime`: the time in milliseconds (a `recv_time`) after which the transaction should be ignored by the trigger service. cannot be more than 100 seconds from the time it is received by the server.

**Note**: for signing you should always use the data type specified in the solidity struct which might be different from the type sent in the request e.g: `recvTime` should be an `uint64` for **Signing** but should be sent as a `string` in the final payload.

## Response

## Success

**Note**: trigger orders can have the following statuses:

* **cancelled**: trigger order was cancelled due to user request, order expiration, or account health issues.
* **triggered**: trigger criteria was met, and order was submitted for execution.
* **internal\_error**: an internal error occurred while processing the trigger order.
* **triggering**: trigger order is currently being processed for execution.
* **waiting\_price**: trigger order is waiting for price criteria to be met.
* **waiting\_dependency**: trigger order is waiting for a dependency order to be filled.
* **twap\_executing**: TWAP order is currently executing individual orders over time.
* **twap\_completed**: TWAP order has completed all scheduled executions.

## Failure


Last updated 1 month ago

---


# List TWAP Executions

Source: https://docs.nado.xyz/developer-resources/api/trigger/queries/list-twap-executions

## Request

Get TWAP executions

`POST [TRIGGER_ENDPOINT]/query`

**Body**

Copy

```
{
  "type": "list_twap_executions",
  "digest": "0x5886d5eee7dc4879c7f8ed1222fdbbc0e3681a14c1e55d7859515898c7bd2038"
}
```

## Request Parameters

Parameter

Type

Required

Description

digest

string

Yes

The digest of the TWAP trigger order to get execution details for.

## Response

## Success

Copy

```
{
  "status": "success",
  "data": {
    "executions": [
      {
        "execution_id": 1,
        "scheduled_time": 1688768157,
        "status": "pending",
        "updated_at": 1688768157050
      },
      {
        "execution_id": 2,
        "scheduled_time": 1688768187,
        "status": {
          "executed": {
            "executed_time": 1688768187050,
            "execute_response": {
              "status": "success",
              "data": {
                "digest": "0x..."
              },
              "id": 12345,
              "request_type": "place_order"
            }
          }
        },
        "updated_at": 1688768187050
      },
      {
        "execution_id": 3,
        "scheduled_time": 1688768217,
        "status": {
          "failed": "Insufficient balance"
        },
        "updated_at": 1688768217050
      },
      {
        "execution_id": 4,
        "scheduled_time": 1688768247,
        "status": {
          "cancelled": "user_requested"
        },
        "updated_at": 1688768247050
      }
    ]
  },
  "request_type": "query_list_twap_executions"
}
```

**Note**: TWAP executions can have the following statuses:

* **pending**: execution is scheduled but has not yet been attempted.
* **executed**: execution was successful, includes execution time and response details from the engine.
* **failed**: execution failed, includes error message.
* **cancelled**: execution was cancelled, includes cancellation reason (e.g., "user\_requested", "linked\_signer\_changed", "expired", "account\_health", "isolated\_subaccount\_closed", "dependent\_order\_cancelled").

## Failure


Last updated 1 month ago

---


# V2

Source: https://docs.nado.xyz/developer-resources/api/v2

Nado V2 API offers REST-based endpoints focused on two key functionalities:

1. **Gateway Queries**: Access real-time market data, including trading pairs and book liquidity.
2. **Archive (indexer) Queries**: Query historical market data, including 24-hour statistics for all products and recent trades.

## Endpoints

## Testnet

* **Gateway**: `https://gateway.test.nado.xyz/v2`
* **Archive (indexer)**: `https://archive.test.nado.xyz/v2`

## Gateway

[Assets](/developer-resources/api/v2/assets)
## Archive (indexer)

[Tickers](/developer-resources/api/v2/tickers)

Last updated 1 month ago

---


# APR

Source: https://docs.nado.xyz/developer-resources/api/v2/apr

## Request

Get pairs

**GET** `[GATEWAY_V2_ENDPOINT]/apr`

## Response

Copy

```
[
    {
        "name": "USDT0",
        "symbol": "USDT0",
        "product_id": 0,
        "deposit_apr": 6.32465621e-10,
        "borrow_apr": 0.010050173557473174,
        "tvl": 20001010125092.633
    },
    {
        "name": "Wrapped BTC",
        "symbol": "WBTC",
        "product_id": 1,
        "deposit_apr": 8.561123e-12,
        "borrow_apr": 0.010050166480005895,
        "tvl": 1045563178297771.2
    }
]
```

## Response Fields

Field name

Type

Nullable

Description

product\_id

number

No

Internal unique ID of spot / perp product

name

string

No

Asset name (as represented internally in the exchange).

symbol

string

No

Asset symbol (as represented internally in the exchange).

deposit\_apr

float

No

The current estimated APR for depositing or holding this asset. **Note**: This value should be multiplied by 100 to represent the percentage (%) form.

borrow\_apr

float

No

The current estimated APR for borrowing this asset. **Note**: This value should be multiplied by 100 to represent the percentage (%) form.

tvl

tvl

No

Total Value Locked (TVL) represents the current USDT0 value of this asset, calculated as the difference between deposits and borrows.


Last updated 1 month ago

---


# Assets

Source: https://docs.nado.xyz/developer-resources/api/v2/assets

## Rate limits

* 1200 requests/min or 20 requests/sec per IP address. (**weight = 2**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Get assets

**GET** `[GATEWAY_V2_ENDPOINT]/assets`

## Response

Copy

```
[
  {
    "product_id": 0,
    "ticker_id": null,
    "market_type": null,
    "name": "USDT0",
    "symbol": "USDT0",
    "taker_fee": null,
    "maker_fee": null,
    "can_withdraw": true,
    "can_deposit": true
  },
  {
    "product_id": 2,
    "ticker_id": "BTC-PERP_USDT0",
    "market_type": "perp",
    "name": "Bitcoin Perp",
    "symbol": "BTC-PERP",
    "maker_fee": 0.0002,
    "taker_fee": 0,
    "can_withdraw": false,
    "can_deposit": false
  },
  {
    "product_id": 1,
    "ticker_id": "BTC_USDT0",
    "market_type": "spot",
    "name": "Bitcoin",
    "symbol": "BTC",
    "taker_fee": 0.0003,
    "maker_fee": 0,
    "can_withdraw": true,
    "can_deposit": true
  }
]
```

## Response Fields

Field name

Type

Nullable

Description

product\_id

number

No

Internal unique ID of spot / perp product

name

string

No

Asset name (as represented internally in the exchange).

symbol

string

No

Asset symbol (as represented internally in the exchange).

maker\_fee

decimal

No

Fees charged for placing a market-making order on the book.

taker\_fee

decimal

No

Fees applied when liquidity is removed from the book.

can\_withdraw

boolean

No

Indicates if asset withdrawal is allowed.

can\_deposit

boolean

No

Indicates if asset deposit is allowed.

ticker\_id

string

Yes

Identifier of a ticker with delimiter to separate base/quote. This is `null` for assets without market e.g: `USDT0`

market\_type

string

Yes

Name of market type (`spot` or `perp`) of asset. This is `null` for assets without a market e.g: `USDT0`


Last updated 1 month ago

---


# Orderbook

Source: https://docs.nado.xyz/developer-resources/api/v2/orderbook

## Rate limits

* 2400 requests/min or 40 requests/sec per IP address. (**weight = 1**)

See more details in [API Rate limits](/developer-resources/api/rate-limits)

## Request

Get orderbook

**GET** `[GATEWAY_V2_ENDPOINT]/orderbook?ticker_id={ticker_id}&depth={depth}`

## Request Parameters

Parameter

Type

Required

Description

ticker\_id

string

Yes

Identifier of a ticker with delimiter to separate base/target.

depth

number

Yes

Number of price levels to retrieve.

## Response

Copy

```
{
    "product_id": 1,
    "ticker_id": "BTC-PERP_USDT0",
    "bids": [
        [
            116215.0,
            0.128
        ],
        [
            116214.0,
            0.172
        ]
    ],
    "asks": [
        [
            116225.0,
            0.043
        ],
        [
            116226.0,
            0.172
        ]
    ],
    "timestamp": 1757913317944
}
```

## Response Fields

Field Name

Type

Nullable

Description

product\_id

u32

No

Unique identifier for the product.

ticker\_id

string

No

Identifier of a ticker with delimiter to separate base/target.

bids

decimal[]

No

An array containing 2 elements. The offer price (first element) and quantity for each bid order (second element).

asks

decimal[]

No

An array containing 2 elements. The ask price (first element) and quantity for each ask order (second element).

timestamp

integer

No

Unix timestamp in milliseconds for when the last updated time occurred.


Last updated 1 month ago

---


# Pairs

Source: https://docs.nado.xyz/developer-resources/api/v2/pairs

## Request

Get pairs

**GET** `[GATEWAY_V2_ENDPOINT]/pairs?market={spot|perp}`

## Request Parameters

Parameter

Type

Required

Description

market

string

No

Indicates the corresponding market to fetch trading pairs for. Allowed values are: `spot`and `perp`. When no `market` param is provided, it returns all available pairs.

## Response

Copy

```
[
    {
        "product_id": 1,
        "ticker_id": "BTC-PERP_USDT0",
        "base": "BTC-PERP",
        "quote": "USDT0"
    },
    {
        "product_id": 2,
        "ticker_id": "ETH-PERP_USDT0",
        "base": "ETH-PERP",
        "quote": "USDT0"
    },
    {
        "product_id": 3,
        "ticker_id": "BTC_USDT0",
        "base": "BTC",
        "quote": "USDT0"
    },
    {
        "product_id": 4,
        "ticker_id": "ETH_USDT0",
        "base": "ETH",
        "quote": "USDT0"
    }
]
```

## Response Fields

Field name

Type

Nullable

Description

product\_id

u32

No

Unique identifier for the product.

ticker\_id

string

No

Identifier of a ticker with delimiter to separate base/target.

base

string

No

Symbol of the base asset.

quote

string

No

Symbol of the target asset.


Last updated 1 month ago

---


# Tickers

Source: https://docs.nado.xyz/developer-resources/api/v2/tickers

## Request

Get pairs

**GET** `[ARCHIVE_V2_ENDPOINT]/tickers?market={spot|perp}&edge={true|false}`

## Request Parameters

Parameter

Type

Required

Description

market

string

No

Indicates the corresponding market to fetch trading tickers info for. Allowed values are: `spot`and `perp`. When no `market` param is provided, it returns all available tickers.

edge

bool

No

Whether to retrieve volume metrics for all chains. When turned off, it only returns metrics for the current chain. Defaults to true.

## Response

**Note**: the response is a map of `ticker_id` -> ticker info object.

Copy

```
{
    "BTC-PERP_USDT0": {
        "product_id": 1,
        "ticker_id": "BTC-PERP_USDT0",
        "base_currency": "BTC",
        "quote_currency": "USDT0",
        "last_price": 25728.0,
        "base_volume": 552.048,
        "quote_volume": 14238632.207250029,
        "price_change_percent_24h": -0.6348599635253989
    }
}
```

## Response Fields

Field Name

Type

Nullable

Description

product\_id

u32

No

Unique identifier for the product.

ticker\_id

string

No

Identifier of a ticker with delimiter to separate base/target.

base\_currency

string

No

Symbol of the base asset.

quote\_currency

string

No

Symbol of the target asset.

last\_price

decimal

No

Last transacted price of base currency based on given quote currency.

base\_volume

decimal

No

24-hours trading volume for the pair (unit in base)

quote\_volume

decimal

No

24-hours trading volume for the pair (unit in quote/target)

price\_change\_percent\_24h

decimal

No

24-hours % price change of market pair


Last updated 1 month ago

---


# Trades

Source: https://docs.nado.xyz/developer-resources/api/v2/trades

## Request

Get pairs

**GET** `[ARCHIVE_V2_ENDPOINT]/trades?ticker_id=BTC-PERP_USDT0&limit=10&max_trade_id=1000000`

## Request Parameters

Parameter

Type

Required

Description

ticker\_id

string

Yes

Identifier of a ticker with delimiter to separate base/target.

limit

integer

No

Number of historical trades to retrieve. Defaults to 100. Max of 500.

max\_trade\_id

integer

No

Max trade id to include in the result. Use for pagination.

## Response

Copy

```
[
    {
        "product_id": 1,
        "ticker_id": "BTC-PERP_USDT0",
        "trade_id": 6351,
        "price": 112029.5896,
        "base_filled": -0.388,
        "quote_filled": 43467.4807648,
        "timestamp": 1757335618,
        "trade_type": "sell"
    },
    {
        "product_id": 1,
        "ticker_id": "BTC-PERP_USDT0",
        "trade_id": 6350,
        "price": 112032.58899999999,
        "base_filled": -0.179,
        "quote_filled": 20053.833431,
        "timestamp": 1757335618,
        "trade_type": "sell"
    }
]
```

## Response Fields

Field Name

Type

Nullable

Description

product\_id

u32

No

Unique identifier for the product.

ticker\_id

string

No

Identifier of a ticker with delimiter to separate base/target.

trade\_id

integer

No

A unique ID associated with the trade for the currency pair transaction.

price

decimal

No

Trade price of base asset in target currency.

base\_filled

decimal

No

Amount of base volume filled in trade.

quote\_filled

decimal

No

Amount of quote/target volume filled in trade.

timestamp

integer

No

Unix timestamp in seconds for when the transaction occurred.

trade\_type

string

No

Indicates the type of the transaction that was completed ("buy" or "sell").


Last updated 1 month ago

---


# Withdrawing (on-chain)

Source: https://docs.nado.xyz/developer-resources/api/withdrawing-on-chain

You can withdraw collateral from Nado directly on-chain, by submitting a slow-mode transaction via the `Endpoint` contract (see Contracts for addresses).

**Note**:

* This is an alternative to withdrawing collateral via our off-chain sequencer. See [Withdraw Collateral](/developer-resources/api/gateway/executes/withdraw-collateral) for more details.
* Slow mode transactions have a 1 USDT0 fee; as such, an approval of 1 USDT0 is required for the slow mode withdrawal to succeed.

## Steps

1. Assemble the bytes needed for a withdraw collateral transaction by encoding the following struct alongside the transaction type `2`:

Copy

```
struct WithdrawCollateral {
    bytes32 sender;
    uint32 productId;
    uint128 amount;
    uint64 nonce;
}
```

1. Submit the transaction via `submitSlowModeTransaction` on our `Endpoint` contract.

## Example

Copy

```
function withdrawNadoCollateral(address nadoEndpoint, bytes32 sender, uint32 productId, uint128 amount) internal {
    WithdrawCollateral memory withdrawal = new WithdrawCollateral(sender, productId, amount, 0);
    bytes memory tx = abi.encodePacked(2, abi.encode(withdrawal));
    IEndpoint(nadoEndpoint).submitSlowModeTransaction(tx);
}
```

Once the transaction is confirmed, it may take a few seconds for it to make its way into the Nado offchain sequencer and for the withdrawal to be processed.


Last updated 1 month ago

---


# 🚀Get Started

Source: https://docs.nado.xyz/developer-resources/get-started

Welcome to Nado! This guide will help you integrate with Nado's API, whether you're building a trading bot, integrating Nado into your app, or just exploring what's possible.

---

## 🆕 New to Nado?

**Coming from a centralized exchange?** Nado works differently - no API keys, wallet-based authentication, and on-chain settlement. Start with Core Concepts to understand the fundamentals.

## ⭐ [Core Concepts](/developer-resources/get-started/core-concepts) - Start Here

**Read this first.** Understand the fundamental differences between Nado and traditional exchanges:

* 🔐 No API keys - how wallet signatures work
* 👤 Subaccounts and the bytes32 format
* 🔗 Linked signers for delegation
* ✍️ EIP-712 signing basics

This guide will save you hours of confusion. Every concept is explained with working code examples in Python, TypeScript, and Rust.

---

## 📚 Quick Links

## 📖 Essential Reading

* ⚡ [**Quickstart**](/developer-resources/get-started/quickstart) - 5-minute end-to-end example (fastest path!)
* 📖 [**Core Concepts**](/developer-resources/get-started/core-concepts) - Fundamental concepts (deep understanding)
* 💰 [**First Deposit**](/developer-resources/get-started/first-deposit) - Step-by-step guide to fund your account
* 🔗 [**Linked Signers**](/developer-resources/get-started/linked-signers) - Complete guide to 1-Click Trading setup
* 📈 **Coming Soon: First Trade** - Place your first order

## ⚠️ Troubleshooting

* 🔧 [**Common Errors**](/developer-resources/get-started/common-errors) - Troubleshooting signature errors, deposit issues, etc.

---

## 🛤️ Choose Your Path

## ⚡ I Want to Integrate Quickly

Perfect for getting up and running fast:

1. ⚡ Follow the [Quickstart](/developer-resources/get-started/quickstart) guide (5 min) - Fastest path to first trade!
2. 📖 Read [Core Concepts](/developer-resources/get-started/core-concepts) (15 min) - Understand the fundamentals
3. 💰 Learn about [First Deposit](/developer-resources/get-started/first-deposit) (10 min) - Fund your account
4. 🔗 Set up [Linked Signers](/developer-resources/get-started/linked-signers) (10 min) - Enable 1-Click Trading
5. 🔧 Bookmark [Common Errors](/developer-resources/get-started/common-errors) for troubleshooting

---

## 🎓 I Want to Understand Deeply

Perfect for building production systems:

1. 📖 Read [Core Concepts](/developer-resources/get-started/core-concepts)
2. ✍️ Read the API [Signing Guide](/developer-resources/api/gateway/signing)
3. 👤 Read about [Subaccounts & Health](/subaccounts-and-health)
4. 📚 Explore the [API Reference](/developer-resources/api)

---

## 💻 I Prefer Using SDKs

Choose your language and dive in:

* 🐍 [Python SDK](https://nadohq.github.io/nado-python-sdk/index.html) - Full-featured, great for data analysis
* 📘 [TypeScript SDK](/developer-resources/typescript-sdk) - Most complete, best for web integrations
* 🦀 [Rust SDK](https://crates.io/crates/nado-sdk) - High performance, ideal for market makers

---

## 💬 Need Help?

We're here to support you:

* 💬 **Telegram Community**: [Join our Telegram](https://t.me/+whsZJKpiiVwwNjQ0) - Get help from the community and Nado team
* 📧 **Report Issues / Feedback**: [Submit a support ticket](https://nado-90114.zendesk.com/hc/en-us/requests/new?ticket_form_id=52275013155481)

**Stuck on something?** The [Core Concepts](/developer-resources/get-started/core-concepts) guide addresses the most common confusions (authentication, subaccounts, linked signers, signing). If you're hitting errors, start there - it'll save you time.


Last updated 12 days ago

---


# 🔧Common Errors

Source: https://docs.nado.xyz/developer-resources/get-started/common-errors

**You're not alone.** Most errors when integrating with Nado fall into a few categories. This guide shows you exactly how to fix them.

---

## 🚨 Error 2028: "Signature does not match with sender's or linked signer's"

**Most Common Error**

This means Nado couldn't verify your EIP-712 signature. This is the most common integration error.

## What It Means

Your signature verification failed. Nado checked your signature against:

1. Your main wallet's address
2. Your subaccount's linked signer address (if one is set)

And neither matched.

---

## 🔑 Cause 1: Wrong Private Key

You're signing with a key that doesn't match either:

* Your main wallet's private key, OR
* Your current linked signer's private key

**How to fix:**

Python

TypeScript

---

## 🌐 Cause 2: Wrong Chain ID

Using the wrong network identifier in your EIP-712 domain.

**Correct chain IDs:**

* **Mainnet**: `57073`
* **Testnet**: `763373`

**How to fix:**

Python

TypeScript

---

## 📝 Cause 3: Wrong Verifying Contract

The `verifyingContract` field in your EIP-712 domain must match the operation type.

**The rules:**

* **For** `place_order`: Use `address(productId)` (e.g., product 2 → `0x0000000000000000000000000000000000000002`)
* **For all other executes**: Use the endpoint contract address

**This is a common gotcha!** Order placement uses a different verifying contract than other operations.

**How to fix:**

Python

TypeScript

**If you're building raw API calls:** Query `/v1/query?type=contracts` to get the endpoint address. For orders, use `address(productId)` as the verifying contract. For other executes, use the endpoint address.

---

## 🔢 Cause 4: Wrong Data Types

EIP-712 requires exact type matching. Common mistakes:

* `expiration` must be `uint64` (number), not string
* `nonce` must be `uint64` (number), not string
* `amount` must be `int128`, not string

**How to fix:**

The SDKs handle type conversion automatically. If you're building raw signatures:

---

## 🐛 Debug Checklist

If you're still getting signature errors, check these in order:

1. **Verify your key**: Print the address of the key you're signing with
2. **Check linked signer**: Query `linked_signer` endpoint to see what's currently set
3. **Verify chain ID**: Mainnet = 57073, Testnet = 763373
4. **Check verifying contract**: Order placement vs other executes use different addresses
5. **Print EIP-712 payload**: Before signing, log the entire typed data structure
6. **Use SDK examples**: The SDKs are battle-tested - compare your code to SDK source

**Pro tip**: If signature works in the UI but fails via API, your linked signer is likely different. Query the linked signer endpoint to verify.

---

## 💰 Error 2024: "Provided address has no previous deposits"

**What It Means**

Your subaccount doesn't exist yet. Subaccounts are created on first deposit (minimum $5 USDT0).

## Common Causes

---

## 📭 Cause 1: No Initial Deposit

You haven't made a deposit yet, or it was less than $5 USDT0.

**How to fix:**

1. Deposit at least $5 USDT0 to create the subaccount
2. Wait for blockchain confirmation (~few seconds on Ink)
3. Verify subaccount exists:

Python

TypeScript

**See also:** [First Deposit Guide](/developer-resources/get-started/first-deposit) for step-by-step deposit instructions

---

## 🌐 Cause 2: Wrong Network

You're querying the wrong network (testnet vs mainnet).

**Symptoms:**

* Subaccount shows balance in UI
* But API returns `exists: false`

**How to fix:**

Verify you're using the correct endpoint:

* **Mainnet**: `gateway.prod.nado.xyz` (chain ID 57073)
* **Testnet**: `gateway.test.nado.xyz` (chain ID 763373)

Python

TypeScript

---

## 📝 Cause 3: Wrong Subaccount bytes32

Your subaccount bytes32 format is incorrect.

**Format:** `bytes32 = wallet address (20 bytes) + subaccount name (12 bytes)`

**How to fix:**

Python

TypeScript

---

## 🔄 Error: 1-Click Trading Broken After API Integration

**Cause: Linked Signer Mismatch**

You updated the linked signer via API, but the UI is still trying to use the old key.

## How to Fix

**Option 1: Reset in UI** (Easiest)

1. Go to Nado UI at [app.nado.xyz](https://app.nado.xyz)
2. Disable 1-Click Trading
3. Re-enable 1-Click Trading (UI will generate new key and link it)

**Option 2: Use Admin Tools**

1. Go to [app.nado.xyz/admin-tools](https://app.nado.xyz/admin-tools)
2. Find "Reset Linked Signer" section
3. Follow instructions to sync API and UI

**Option 3: Update UI to Use Your Key**

If you want the UI to use YOUR linked signer (not a random one):

1. Query your current linked signer:

Python

TypeScript

1. In UI: Manually enter that private key when enabling 1-Click Trading

---

## 📡 Error: Wrong Endpoint / Network Mismatch

## Symptoms

* API returns errors but UI works fine
* "Subaccount not found" but you see balance in UI
* Signature works in UI but fails via API

## Solution

Verify you're using matching endpoints:

Environment

Gateway URL

Chain ID

Use When

**Mainnet**

`gateway.prod.nado.xyz`

57073

Real funds, production

**Testnet**

`gateway.test.nado.xyz`

763373

Testing, development

Check what the UI is using:

1. Go to [app.nado.xyz](https://app.nado.xyz)
2. Look at network indicator (top right)
3. Match your API endpoint to that network

---

## 💡 General Debugging Tips

## Start with SDK Examples

The SDKs are battle-tested. If you're hitting errors:

1. Find a similar example in the SDK docs:

   * [Python SDK](https://nadohq.github.io/nado-python-sdk/index.html)
   * [TypeScript SDK](/developer-resources/typescript-sdk)
   * [Rust SDK](https://crates.io/crates/nado-sdk)
2. Compare your code to the example
3. Use the SDK method directly (let it handle the complexity)

---

## Enable Debug Logging

Python

TypeScript

Rust

---

## Verify with Raw API Call

If SDK is failing, try a raw curl to isolate the issue:

If raw API works but SDK fails → SDK configuration issue If raw API also fails → Network/endpoint issue

---

## 🆘 Still Stuck?

If you've tried everything above and still hitting errors:

## 1. Check What's Different

What works:

* UI? (means your wallet/funds are correct)
* SDK examples? (means SDK is configured correctly)
* Raw curl? (means API is accessible)

What fails:

* Your specific code? (code issue)
* All SDK methods? (configuration issue)
* Everything including curl? (network/endpoint issue)

## 2. Get Help

* **Telegram Community**: [Join our Telegram](https://t.me/+whsZJKpiiVwwNjQ0) - Paste your error, we'll help debug
* **Submit a ticket**: [Nado Support](https://nado-90114.zendesk.com/hc/en-us/requests/new?ticket_form_id=52275013155481)

When asking for help, include:

* Error message (exact text)
* Network (mainnet/testnet)
* SDK and version
* What you're trying to do (place order, deposit, etc.)
* Code snippet (remove private keys!)

---

**Remember:** Most integration errors are one of the 6 issues above. Check the list systematically and you'll find the culprit.

---

## Next Steps

* **Back to basics**: [Core Concepts](/developer-resources/get-started/core-concepts) - Review fundamentals
* **Deep dive**: [Linked Signers Guide](/developer-resources/get-started/linked-signers) - Master the most confusing part
* **Get trading**: [First Trade Guide](https://github.com/nadohq/nado-docs/blob/main/docs/developer-resources/get-started/first-trade.md) - Place your first order


Last updated 12 days ago

---


# 📖Core Concepts

Source: https://docs.nado.xyz/developer-resources/get-started/core-concepts

If you're coming from a centralized exchange (CEX) API, Nado works differently. This guide explains the fundamental concepts that are different from what you might expect, so you can integrate successfully without confusion.

---

## What Makes Nado Different?

Nado is a **decentralized exchange**. This means:

* ❌ **No API keys** - You sign requests with your wallet's private key
* ❌ **No username/password** - Your wallet address is your identity
* ✅ **Full custody** - You control your funds at all times
* ✅ **On-chain settlement** - Trades settle on the blockchain

If these concepts are new to you, don't worry. This guide will walk you through everything step by step.

---

## 1. Authentication: Wallet Signatures, Not API Keys

**TL;DR:** Nado uses wallet signatures (EIP-712) instead of API keys. Your wallet's private key IS your authentication. Executes (writes) need signatures; queries (reads) don't.

---

## 🔑 The CEX Model (What You Might Expect)

On centralized exchanges like Binance or Coinbase:

1. You create an account → get API key + secret
2. You sign requests with HMAC using that secret
3. The exchange validates your API key

---

## 🔐 The Nado Model (How It Actually Works)

On Nado:

1. You have a wallet (like MetaMask, a hardware wallet, or just a private key)
2. You sign requests with your wallet's private key using **EIP-712** (EVM signing standard)
3. Nado validates your cryptographic signature

**Why no API keys?** Because Nado is decentralized - there's no central server to issue API keys. Your wallet's private key IS your authentication.

---

## 📝 What This Means for You

**Every write operation (execute) must be cryptographically signed.**

This includes operations like:

* Placing orders
* Canceling orders
* Depositing/withdrawing funds
* Linking signers

**Good news:** Read operations (queries) don't require signing! Things like checking balances, getting market data, or querying order status work without signatures.

Here's a simple example of setting up a client that handles signing for you:

Python

TypeScript

**🔒 Security Critical:** Your private key gives FULL access to your funds. Never share it, commit it to git, or send it over unencrypted channels.

---

## ✍️ Two Ways to Sign Executes

When performing write operations (executes), you can sign with **either**:

1. **Your main wallet's private key** (always works)
2. **A linked signer's private key** (optional - explained in Section 3)

**Common Issue:** Signature errors often happen when users don't realize they can use either key. Both are cryptographically valid!

---

## 2. Subaccounts: Your Trading Compartments

**TL;DR:** Subaccounts are trading compartments within your wallet. Format = wallet address (20 bytes) + name (12 bytes). Must deposit ≥$5 USDT0 to create. Start with "default" subaccount.

---

## 📦 What is a Subaccount?

A subaccount is a **trading compartment** within your wallet. Think of it like having multiple trading accounts tied to the same wallet address.

**Format**: Subaccounts are identified by a `bytes32` value that combines:

* **20 bytes**: Your wallet address (e.g., `0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb`)
* **12 bytes**: Subaccount name (e.g., `default`, `strategy1`, etc.)

**Example**:

---

## 🎯 The Default Subaccount

When you first interact with Nado, you'll use the **"default" subaccount**. This is standard across Nado:

* UI users: Automatically use "default"
* API users: Should use "default" to start

Python

TypeScript

---

## 💰 Creating a Subaccount: The Critical Step

**⚠️ CRITICAL**: A subaccount doesn't exist until you make an initial deposit of **at least $5 USDT0** (or equivalent).

This trips up almost everyone. Here's what happens:

**Before deposit**:

**After deposit (≥ $5)**:

**Why?** Nado is on-chain. Creating a subaccount requires an on-chain transaction, which happens when you make your first deposit.

**Common Error**: Trying to place an order before depositing → `"error": "The provided address has no previous deposits"`

**Solution**: Make an initial deposit first. See our [First Deposit](https://github.com/nadohq/nado-docs/blob/main/docs/developer-resources/get-started/first-deposit.md) guide.

---

## 🔢 Multiple Subaccounts

* **UI limit**: 4 subaccounts (`default`, `default_1`, `default_2`, `default_3`)
* **API limit**: Unlimited! You can create as many as you want programmatically

**Each subaccount is completely independent:**

* Independent balances and positions
* Independent health (risk) calculations
* Can be liquidated independently (one doesn't affect others)
* Can have its own linked signer (see Section 3)

**Common use cases:**

* 🎯 Separate strategies (one for spot, one for perps)
* 📊 Different risk levels (conservative vs aggressive)
* 👥 Team management (each trader gets a subaccount)
* 🧪 Testing (use a separate subaccount to test new strategies)

---

## 3. Linked Signers: Delegation for Security

**TL;DR:** Linked signers let you authorize a separate key to trade on your behalf. Optional but recommended for security. Main wallet = master key, linked signer = trading-only key. Can revoke anytime.

**Skip this if you're just getting started.** Linked signers are optional. You can always sign with your main wallet's private key.

Come back here when you see "1-Click Trading" or "1CT" mentioned, or when you're ready to set up delegated signing for security.

---

## 🔗 What is a Linked Signer?

A **linked signer** is a separate private key that you authorize to trade on behalf of your subaccount.

**Think of it like this**:

* 🔑 **Your main wallet** = Your master key (controls everything, including funds)
* ✍️ **Linked signer** = A trading-only key (can place orders, cancel orders, but you control the authorization)

---

## 🛡️ Why Would You Use One?

**Security**: Keep your main wallet private key offline/secure, use a linked signer for day-to-day trading.

**Common scenario**:

1. Main wallet: Hardware wallet (very secure, inconvenient for API trading)
2. Linked signer: Hot wallet private key (less secure, convenient for API)
3. If linked signer is compromised → disable it, main wallet is still safe

---

## ⚙️ How It Works

1. **Generate a new private key** (or use an existing one):

Python

TypeScript

Rust

---

1. **Link it to your subaccount** (must sign with main wallet):

Python

TypeScript

---

1. **To revoke/unlink a signer** (set back to empty):

Python

TypeScript

---

1. **Now you can sign executes with either**:

   * Main wallet private key (always works!)
   * Linked signer private key (if one is set)

---

## ✅ Verifying Your Current Linked Signer

**Common Issue**: "Why is my signature failing?" → Often because you're using the wrong private key.

**Solution**: Check which linked signer is currently set for your subaccount.

Python

TypeScript

---

## 🖱️ UI 1-Click Trading (1CT)

If you see "1-Click Trading" or "1CT" in the Nado UI:

* This is a linked signer that the UI creates for you
* It generates a random private key and stores it locally
* It makes trading faster (no wallet popup for every order)

**Important for API users:**

* If you set a linked signer via API, the UI's 1CT key becomes invalid
* If the UI breaks after you change the linked signer via API:

  1. Go to UI settings
  2. Disable 1-Click Trading
  3. Re-enable it with your new linked signer's private key (or let it generate a new one)

For a complete guide on linked signers, see our [Linked Signers Deep Dive](/developer-resources/get-started/linked-signers).

---

## 4. Signing Requests: EIP-712

**TL;DR:** EIP-712 is the EVM signing standard. Executes (writes) need EIP-712 signatures; queries (reads) don't. SDKs handle this automatically.

---

## 🔐 What is EIP-712?

EIP-712 is the EVM standard for signing structured data. Every execute (write operation) you make to Nado—like placing orders, canceling orders, depositing, withdrawing—must be signed using EIP-712.

**Queries (reads) don't need signing!** Operations like checking balances or getting market data work without signatures.

**Good news**: The SDKs handle all of this automatically. You don't need to understand the details unless you're implementing raw API calls.

---

## ⚙️ The Signing Process (Behind the Scenes)

When you call `client.place_order(...)`, here's what happens:

1. **SDK creates structured data**:

1. **SDK signs it with your private key** using EIP-712
2. **SDK sends request** with the signature
3. **Nado verifies** the signature matches your wallet or linked signer

## ⚠️ Common Signing Errors

**Error: "Signature does not match with sender's or linked signer's"**

This means Nado couldn't verify your signature. This is the most common error when getting started.

---

**🔑 Cause 1: Wrong Private Key**

You're signing with a key that isn't:

* Your main wallet's private key, OR
* Your linked signer's private key

**Solution**: Verify which linked signer is set (see Section 3 above), then use the correct key.

---

**🌐 Cause 2: Wrong Chain ID**

Using the wrong network identifier:

* **Mainnet**: `57073`
* **Testnet**: `763373`

**Solution**: Make sure you're using the correct network when creating your client.

---

**📝 Cause 3: Wrong Verifying Contract**

The `verifyingContract` field in EIP-712 must match the operation:

* For `place_order`: Use `address(productId)` (e.g., product 2 → `0x0000000000000000000000000000000000000002`)
* For other operations: Use the endpoint contract address

**Solution**: The SDK handles this automatically. If you're making raw API calls, check the [Signing Reference](/developer-resources/api/gateway/signing).

---

**🔢 Cause 4: Wrong Data Types**

Type mismatches in the signature payload:

* `expiration` must be `uint64` (number), not string
* `nonce` must be `uint64` (number), not string

**Solution**: The SDK handles this. If implementing manually, check the [EIP-712 Examples](/developer-resources/api/gateway/signing/examples).

---

For detailed troubleshooting and more error scenarios, see our [Common Errors Guide](/developer-resources/get-started/common-errors).

---

## 📋 Quick Reference

## 🔐 Authentication

* ❌ No API keys
* ✅ Executes (writes) require wallet signature (EIP-712)
* ✅ Queries (reads) don't require signatures
* ✅ Can sign executes with main wallet OR linked signer

## 👤 Subaccounts

* Format: `bytes32` = wallet address (20 bytes) + name (12 bytes)
* Default subaccount: `"default"`
* Must deposit ≥ $5 to create
* Check if exists: `client.get_subaccount_info("default")` → look for `"exists": true`

## 🔗 Linked Signers

* Optional delegation mechanism for signing executes
* One per subaccount
* Can sign executes with EITHER main wallet OR linked signer
* Verify current: `client.get_linked_signer("default")`
* UI "1-Click Trading" = linked signer

## ✍️ Signing

* All executes (writes) use EIP-712 signatures
* Queries (reads) don't require signatures
* SDKs handle signing automatically
* Mainnet chain ID: `57073`
* Testnet chain ID: `763373`

---

## Next Steps

Now that you understand the core concepts, you're ready to:

1. [**Make Your First Deposit**](https://github.com/nadohq/nado-docs/blob/main/docs/developer-resources/get-started/first-deposit.md) - Fund your subaccount
2. [**Place Your First Order**](https://github.com/nadohq/nado-docs/blob/main/docs/developer-resources/get-started/first-trade.md) - Start trading
3. [**Set Up a Linked Signer**](/developer-resources/get-started/linked-signers) - Optional: For better security

If you run into issues, check our [Common Errors](/developer-resources/get-started/common-errors) troubleshooting guide.

---

**Remember**: If you're confused, you're not alone. These concepts are different from traditional exchanges. Take your time, and don't hesitate to ask for help in our [Telegram community](https://t.me/+whsZJKpiiVwwNjQ0) or [report an issue](https://nado-90114.zendesk.com/hc/en-us/requests/new?ticket_form_id=52275013155481).


Last updated 17 days ago

---


# 💰First Deposit

Source: https://docs.nado.xyz/developer-resources/get-started/first-deposit

**TL;DR**: You need >= $5 USDT0 to create a subaccount. Three methods: (1) **UI Deposit** (easiest - start here!), (2) **On-Chain Contract Call** (recommended for developers), or (3) **Direct Deposit** (alternative, requires caution). We recommend starting on testnet with free faucet tokens!

---

## 📋 Prerequisites

Before depositing, you'll need:

* A wallet with a private key
* Funds to deposit:

  + **Testnet**: Free tokens from [Nado Testnet Faucet](https://testnet.nado.xyz/portfolio/faucet)
  + **Mainnet**: >= $5 USDT0 or equivalent
* Gas for transaction fees:

  + **Testnet**: Free Ink Sepolia ETH from [Ink Faucet](https://docs.inkonchain.com/tools/faucets)
  + **Mainnet**: Ink ETH

**Starting on testnet?** You can get both testnet USDT0 (for trading) and Ink Sepolia ETH (for gas) completely free from the faucets above. This is the recommended way to learn!

---

## 🎯 Why You Need to Deposit

On Nado, **subaccounts are created automatically when you make your first deposit** (>= $5 USDT0). Unlike centralized exchanges, there's no separate "account creation" step - depositing funds creates your subaccount.

**Important**: Subaccounts don't exist until you deposit. If you try to query a subaccount before depositing, you'll get `"exists": false`.

---

## 🚀 Method 1: UI Deposit (Easiest)

This is the **recommended method for getting started** - use the Nado web interface. Perfect for your first deposit!

**For Testnet (Ink Sepolia):**

1. Go to [Nado Testnet Faucet](https://testnet.nado.xyz/portfolio/faucet)
2. Connect your wallet (make sure you're on Ink Sepolia network)
3. Mint testnet USDT0 (click "Mint" button - free!)
4. Go to [Testnet Trading](https://testnet.nado.xyz) and verify your balance
5. You're ready to trade! (subaccount created automatically)

**For Mainnet (Ink):**

1. Go to [Nado Mainnet](https://nado.xyz)
2. Connect your wallet (make sure you're on Ink Mainnet network)
3. Click "Deposit" and follow the UI instructions
4. Send >= $5 USDT0 to complete deposit
5. Subaccount created automatically

**Why start here?** The UI handles all the complexity - approvals, contract calls, and verification. This is the safest method.

---

## ⚙️ Method 2: On-Chain Contract Call (Recommended for Developers)

This is the **recommended method for programmatic deposits**. Call the Endpoint contract directly to deposit funds.

## Step 1: Get Contract Addresses

Python

TypeScript

Rust

---

## Step 2: Approve Token Allowance and Deposit

Python

TypeScript

Rust

---

## Step 3: Verify Your Deposit

Wait ~30 seconds for blockchain confirmation, then check if your subaccount was created:

Python

TypeScript

Rust

**Success!** If `exists: true`, your subaccount is ready and you can start trading! 🎉

---

## 🔄 Method 3: Direct Deposit (Alternative)

Each subaccount has a unique deposit address. You can send funds directly to this address and they will automatically be credited.

**⚠️ Use with Caution**: This method is the easiest but requires extreme care. **Only send supported tokens** - sending unsupported tokens will result in permanent loss of funds. We recommend using Method 1 (UI) or Method 2 (On-Chain) instead.

## Step 1: Get Your Direct Deposit Address

Python

TypeScript

Rust

---

## Step 2: Verify Supported Tokens (CRITICAL)

Before sending any funds, **you MUST verify the token is supported**:

Python

TypeScript

Rust

---

## Step 3: Send Funds Directly (From Any Wallet)

**That's it!** Just send the supported token to your direct deposit address:

* **From MetaMask**: Send transaction to the deposit address
* **From CEX**: Withdraw to the deposit address (use correct network!)
* **From another wallet**: Transfer tokens to the deposit address

**The funds will be automatically credited to your Nado subaccount** within ~30 seconds.

**CRITICAL WARNINGS**:

* ⚠️ **ONLY send supported tokens** (verify with Step 2 above)
* ⚠️ **Use the correct network**: Ink Mainnet or Ink Sepolia Testnet
* ⚠️ **Sending wrong tokens = permanent loss** - no refunds possible
* ⚠️ **Double-check before sending** - you cannot undo this

---

## Step 4: Verify Your Deposit

Use the same verification code from Method 2, Step 4 to confirm your subaccount was created and has the deposited balance.

---

## 🎁 Testnet Faucet

**New to Nado?** Start on testnet with free tokens!

1. **Get Free Testnet USDT0**: <https://testnet.nado.xyz/portfolio/faucet>
2. **Get Free Ink Sepolia ETH** (for gas): [Ink Faucet](https://docs.inkonchain.com/tools/faucets)

---

## 🤔 Which Method Should I Use?

Method

Best For

Pros

Cons

**UI Deposit**

First-time users, getting started quickly

Safest, handles everything

Requires web interface

**On-Chain Contract**

Developers, bots, programmatic deposits

Type-safe, explicit, fails early

Requires code

**Direct Deposit**

CEX withdrawals, simple transfers

No contract calls, works with any wallet

Risk of sending wrong token

**Recommendation**: Start with **Method 1 (UI)** to get familiar, then use **Method 2 (On-Chain)** for your production bots and applications.

---

## 🆘 Troubleshooting

## Subaccount not created

**Possible causes:**

1. Deposit amount < $5 USDT0 minimum
2. Transaction still confirming (wait 30-60 seconds)
3. Wrong token sent (if using direct deposit)

**Debug steps:**

1. Check transaction status on block explorer
2. Verify balance with the verification code above
3. Check [Common Errors](/developer-resources/get-started/common-errors) for help

## Transaction failed

**Possible causes:**

1. Insufficient gas
2. Token allowance not approved (Method 2 only)
3. Wrong product ID or amount

**Solution**: Review error message and check your parameters

---

## 🎯 Next Steps

Now that you have funds deposited, you're ready to trade!

* 📈 [**Place Your First Order**](/developer-resources/api/gateway/executes/place-order) - Start trading
* 🔗 [**Set Up Linked Signers**](/developer-resources/get-started/linked-signers) - Enable 1-Click Trading
* 📖 [**Read Core Concepts**](/developer-resources/get-started/core-concepts) - Understand how Nado works

**Success!** You've completed your first deposit. Happy trading! 🎉


Last updated 12 days ago

---


# 🔗Linked Signers

Source: https://docs.nado.xyz/developer-resources/get-started/linked-signers

**TL;DR:** Linked signers let you delegate trading permissions to a separate private key, keeping your main wallet secure. This is Nado's "1-Click Trading" feature.

**This guide addresses the most common questions about linked signers.**

---

## 🤔 What is a Linked Signer?

A **linked signer** is a separate private key that you authorize to sign trading operations on behalf of your subaccount.

Think of it like this:

* **Main wallet**: Your master key (high value, kept secure)
* **Linked signer**: Delegated key (trading only, can be rotated)

## Key Facts

* **One per subaccount**: Each subaccount can have exactly one linked signer
* **Full permissions**: Linked signer can do ANYTHING your main wallet can (trade, withdraw, etc.)
* **Optional**: You don't need a linked signer - your main wallet always works
* **UI's "1-Click Trading"**: This is just a linked signer with a random key

**Important**: When you have a linked signer set, you can sign requests with EITHER your main wallet key OR the linked signer key. Both work.

---

## 🎯 Why Use a Linked Signer?

## 🔐 Security

Keep your main wallet (with all your assets) offline or in cold storage. Use a linked signer key only for trading on Nado.

**Example**: Your main wallet has $1M in assets. Create a linked signer for trading bots with only $10K exposure on Nado. If the bot key is compromised, attacker can only access your Nado subaccount (limited by what's deposited).

---

## 🤖 Programmatic Trading

Run trading bots without exposing your main wallet private key to servers.

**Example**: Deploy a market-making bot to a cloud server. Use a linked signer key instead of your main wallet key. If the server is compromised, you can revoke the linked signer and your main wallet remains safe.

---

## ⚡ Convenience (1-Click Trading)

The Nado UI's "1-Click Trading" feature generates a random linked signer key and stores it encrypted in your browser. This means:

* No MetaMask popup for every trade
* Faster execution
* Same security (key never leaves your device)

---

## 🛠️ How to Set Up a Linked Signer

**Before you start**: Your subaccount must exist (requires initial deposit >= $5 USDT0). See [Core Concepts](/developer-resources/get-started/core-concepts) if you haven't created a subaccount yet.

## Method 1: Via API (Recommended for Developers)

## Step 1: Generate a New Private Key

Python

TypeScript

Rust

---

## Step 2: Link the Signer to Your Subaccount

**Critical**: This request must be signed with your **MAIN WALLET** private key, not the linked signer key!

Python

TypeScript

---

## Step 3: Verify It Worked

Python

TypeScript

---

## Method 2: Via UI (1-Click Trading)

This is the easiest method if you're just getting started:

1. **Go to Nado UI**: [app.nado.xyz](https://app.nado.xyz)
2. **Connect your wallet** (MetaMask, WalletConnect, etc.)
3. **Open Settings** (gear icon)
4. **Find "1-Click Trading"** section
5. **Click "Enable"**

The UI will:

* Generate a random private key
* Link it to your default subaccount
* Store it encrypted in your browser
* You're done!

**UI-generated vs API**: If you enable 1-Click Trading in the UI, then later link a different signer via API, the UI's 1-Click Trading will break (it'll be using the old key). See [troubleshooting below](/developer-resources/get-started/linked-signers#error-ui-1-click-trading-broken).

---

## 🔍 How to Check Your Current Linked Signer

## Via API

Python

TypeScript

## Via Raw API Call

Response:

---

## ✍️ How to Sign Requests with Your Linked Signer

Once you've linked a signer, you have two options for signing trading requests:

## Option 1: Use Main Wallet Key (Always Works)

Python

TypeScript

---

## Option 2: Use Linked Signer Key (Recommended for Trading Bots)

**Critical Understanding**: When using a linked signer, the `sender` field is **ALWAYS your main wallet's subaccount**, but the signature comes from the **linked signer's private key**. This is the core concept of linked signers.

Python

TypeScript

**How linked signers work**:

1. **Sender field**: Always your main wallet's subaccount (the account that owns the funds)
2. **Signature**: Created by the linked signer's private key
3. **Verification**: Nado checks if the signature matches either the main wallet OR the linked signer for that subaccount
4. **Result**: Request is accepted if the linked signer is authorized for that subaccount

---

## 🏗️ Advanced: On-Chain Linked Signer Setup (Smart Contracts & Safe Wallets)

**Use Cases**: This advanced setup method enables:

* **Smart contract integration**: Your smart contract can trade on Nado by linking an EOA signer
* **Safe (Gnosis Safe) wallet trading**: Multi-sig Safe wallets can link an EOA for trading
* **Contract-to-contract**: Non-EOA accounts (contracts) as the main account, EOA as the trading signer

## Why Use On-Chain Setup?

The standard linked signer setup (via API) requires an EOA (Externally Owned Account) with a private key to sign the link request. But what if:

* Your main account is a **smart contract** (not an EOA)?
* You want to use a **Safe wallet** (multi-sig contract) to trade?
* You're building a **protocol integration** where the main account is a contract?

**Solution**: Use **slow mode transactions** to link a signer directly on-chain. This bypasses the EIP-712 signature requirement and allows contracts to link signers.

---

## How It Works

1. **Your smart contract** calls `submitSlowModeTransaction` on the Nado Endpoint contract
2. **Transaction includes**: LinkSigner data (sender = contract address, signer = EOA address)
3. **Nado sequencer** picks up the on-chain transaction
4. **Result**: The EOA is now authorized to sign trades for the contract's subaccount

---

## Setup Steps

## Prerequisites

* Smart contract deployed on-chain (or Safe wallet)
* ≥ 1 USDT0 for slow mode fee
* EOA private key to use as the linked signer

## Step 1: Deposit Funds to Contract Subaccount

Your contract must first deposit into Nado to create its subaccount:

## Step 2: Link Signer via Slow Mode Transaction

Solidity (Smart Contract)

Safe Wallet (via UI)

Python (Contract Integration)

**After submission:**

* Wait ~5-10 seconds for sequencer to process
* EOA can now sign trades for the contract's subaccount

**For Safe (Gnosis Safe) wallets:**

1. **Generate an EOA** to use as linked signer:
2. **Option A: Use Nado UI**

   * Connect your Safe wallet to [app.nado.xyz](https://app.nado.xyz)
   * Go to Settings → 1-Click Trading
   * Enable 1-Click Trading
   * **Paste the EOA private key** (from step 1) instead of generating random
   * Safe will propose a transaction to owners for approval
   * Once approved and executed, EOA is linked
3. **Option B: Use Safe Transaction Builder**

   * Go to Safe's Transaction Builder app
   * Call `submitSlowModeTransaction` on Nado Endpoint
   * Use the Solidity code above as reference for transaction data
   * Owners approve and execute

---

## Important Notes

**Slow Mode Fee**: Requires 1 USDT0 fee for each slow mode transaction. Make sure your contract has approved the Nado Endpoint for this amount.

**Processing Time**: Slow mode transactions take ~5-10 seconds to be picked up by the Nado sequencer. Query the linked signer endpoint to verify it's been processed.

**Use Cases**:

* **DeFi protocols**: Your lending protocol's contract can trade on Nado
* **DAOs**: DAO treasury (multi-sig) can link a trading bot EOA
* **Automated strategies**: Contract-based strategies can execute trades via linked EOA
* **Safe wallets**: Use Safe for custody, EOA for trading (best of both worlds)

---

## Verification

After submitting the slow mode transaction, verify the link worked:

Expected response:

---

## Complete Example: Safe Wallet Trading

**Scenario**: You have a Safe wallet with $100K and want to trade on Nado without exposing the Safe's keys.

**Setup:**

1. Generate a random EOA (linked signer)
2. Deposit $100K from Safe to Nado (creates Safe's subaccount)
3. Use Nado UI or Safe Transaction Builder to link the EOA
4. Run a trading bot with the EOA's private key
5. Bot places trades for the Safe's subaccount
6. Withdrawals go back to the Safe wallet address

**Security**: Even if the EOA key is compromised, attacker can only trade (not withdraw to their own address). Funds always return to the Safe.

---

## Learn More

For complete smart contract integration details, see:

* [**Integrate via Smart Contracts**](/developer-resources/api/integrate-via-smart-contracts) - Full technical documentation
* [**Endpoint Contract Source**](https://github.com/nadohq/nado-contracts/blob/main/core/contracts/Endpoint.sol) - Nado's on-chain contracts

---

## 🔄 How to Update Your Linked Signer

Need to rotate your linked signer key? Here's how:

Python

TypeScript

**Important**: Linking a new signer immediately revokes the old one. There's no grace period. Update your trading bots to use the new key before revoking.

---

## ❌ How to Disable (Revoke) Your Linked Signer

To completely remove the linked signer:

Python

TypeScript

---

## ⚠️ Common Issues & Solutions

## Error: "Signature does not match" (Error 2028)

**Cause**: You're signing with the wrong private key.

**Debug steps:**

1. Check what linked signer is currently set:

1. Verify your signing key's address:

1. Compare:

   * If no linked signer (zero address): Use main wallet key
   * If linked signer is set: Use that key OR main wallet key

**Solution**: Use the correct private key when creating your client.

---

## Error: UI 1-Click Trading Broken

**Symptoms:**

* 1-Click Trading worked before
* You updated linked signer via API
* Now UI shows errors when placing orders

**Cause**: UI is still using the old linked signer key, but you changed it via API.

**Solution Option 1 (Easiest):**

1. Go to UI Settings
2. Disable 1-Click Trading
3. Re-enable 1-Click Trading
4. UI will generate a NEW key and link it

**Solution Option 2 (Keep Your API Key):**

1. Query your current linked signer via API (the one you want to keep)
2. In UI Settings: Disable 1-Click Trading
3. Enable 1-Click Trading with "Use existing key"
4. Paste your API's linked signer private key

**Solution Option 3 (Admin Tools):**

1. Go to [app.nado.xyz/admin-tools](https://app.nado.xyz/admin-tools)
2. Use "Reset Linked Signer" to sync API and UI

---

## Error: Hit Rate Limit When Updating Linked Signer

**Cause**: You've changed linked signer 50+ times in the past 7 days.

**Solution**:

Check your rate limit usage:

Response shows how many operations you have left. Wait for the 7-day window to roll over, or use your current linked signer.

---

## 🔒 Security Considerations

## Linked Signer Has Full Access

**Critical**: A linked signer can do ANYTHING your main wallet can do on your subaccount:

* Place and cancel orders
* Withdraw funds (funds go to the **main wallet**, not the linked signer)
* Change the linked signer
* Mint/burn NLP

Treat the linked signer private key like your main wallet key.

**Important**: While a linked signer can initiate withdrawals, **all withdrawals always go to the main wallet address** (the subaccount owner). The linked signer cannot withdraw funds to its own address or any other address.

---

## Best Practices

1. **Generate unique keys**: Don't reuse private keys across subaccounts or platforms
2. **Rotate regularly**: Change linked signer every 30-90 days
3. **Monitor activity**: Check your subaccount for unexpected activity
4. **Limit exposure**: Only deposit what you need for trading
5. **Secure storage**: Store linked signer keys in environment variables or secrets manager, never in code

---

## Compromised Linked Signer?

If you suspect your linked signer key was compromised:

1. **Immediately revoke** it (set to zero address) using your main wallet
2. **Withdraw funds** from the subaccount if needed
3. **Generate a new** linked signer key
4. **Investigate** how the compromise happened

---

## 📊 Rate Limits

**Limit**: 50 linked signer operations per subaccount per 7-day rolling window

**What counts as an operation:**

* Linking a new signer
* Updating linked signer
* Revoking linked signer (setting to zero)

**Check your usage:**

Bash (curl)

Python

TypeScript

50 operations per week is generous for normal use. If you hit the limit, you're likely testing or have a misconfigured bot updating the signer repeatedly.

---

## 📚 Related Documentation

* [**Core Concepts**](/developer-resources/get-started/core-concepts) - Fundamentals of authentication and subaccounts
* [**Common Errors**](/developer-resources/get-started/common-errors) - Troubleshooting signature errors
* [**API Reference: Link Signer**](/developer-resources/api/gateway/executes/link-signer) - Technical API docs
* [**API Reference: Query Linked Signer**](/developer-resources/api/gateway/queries/linked-signer) - Query endpoint docs

---

## 🆘 Need Help?

Still confused about linked signers?

* **Telegram Community**: [Join our Telegram](https://t.me/+whsZJKpiiVwwNjQ0) - Ask questions, we're here to help!
* **Submit a ticket**: [Nado Support](https://nado-90114.zendesk.com/hc/en-us/requests/new?ticket_form_id=52275013155481)

**You've got this!** Linked signers seem complex at first, but once set up, they make trading much more convenient and secure. Take your time, follow the examples above, and don't hesitate to ask for help.


Last updated 12 days ago

---


# ⚡Quickstart

Source: https://docs.nado.xyz/developer-resources/get-started/quickstart

**Goal:** Get from zero to your first trade on Nado in ~5 minutes.

This guide assumes you have basic blockchain knowledge and want to jump straight in. For deeper understanding, see [Core Concepts](/developer-resources/get-started/core-concepts).

---

## ✅ Prerequisites

Before you start, make sure you have:

* **Ink wallet with private key** (any EVM-compatible wallet works)
* **Ink Sepolia ETH for gas** - Get free testnet ETH from [Ink Faucet](https://docs.inkonchain.com/tools/faucets)
* **Testnet USDT0** (≥ $5) - Mint free tokens from [Nado Testnet Faucet](https://testnet.nado.xyz/portfolio/faucet)

**This guide uses Testnet (Ink Sepolia)** for safe learning without real funds.

**Get testnet tokens:**

1. **Gas (Ink Sepolia ETH)**: [Ink Faucet](https://docs.inkonchain.com/tools/faucets) - Free testnet ETH for transaction fees
2. **Trading funds (USDT0, KBTC, etc.)**: [Nado Testnet Faucet](https://testnet.nado.xyz/portfolio/faucet) - Mint testnet tokens

When ready for real trading, use Mainnet (Ink) and replace `testnet` with `mainnet` in all examples.

---

## 📦 Step 1: Install SDK (2 minutes)

Choose your preferred language and install the SDK:

Python

TypeScript

Rust

**Verify installation:**

**Verify installation:**

**Or manually add to Cargo.toml:**

**Verify installation:**

---

## 🌐 Step 2: Connect to Nado (30 seconds)

Create a client and verify connection:

Python

TypeScript

**Security:** Never commit private keys to git! Use environment variables:

* Python: `os.getenv('PRIVATE_KEY')`
* TypeScript: `process.env.PRIVATE_KEY`
* Rust: `std::env::var("PRIVATE_KEY")`

---

## 💰 Step 3: Create Subaccount (Deposit Funds) (1 minute)

Subaccounts are created automatically on your first deposit. You need **≥ $5 USDT0** to create a subaccount.

**Multiple deposit methods available:**

* **UI Faucet** (easiest for testnet - shown below)
* **Direct Deposit** (send tokens to your unique deposit address)
* **Smart Contract** (on-chain depositCollateral call)

See the [Depositing Guide](/developer-resources/api/depositing) for all methods and advanced options.

## Option A: Deposit via UI (Easiest)

**For Testnet (Ink Sepolia):**

1. Go to [Nado Testnet Faucet](https://testnet.nado.xyz/portfolio/faucet)
2. Connect your wallet (make sure you're on Ink Sepolia network)
3. Mint testnet USDT0 (click "Mint" button - free!)
4. Go to [Testnet Trading](https://testnet.nado.xyz) and verify your balance
5. You're ready to trade! (subaccount created automatically)

**For Mainnet (Ink):**

1. Go to [Nado Mainnet](https://nado.xyz)
2. Connect your wallet (make sure you're on Ink Mainnet network)
3. Click "Deposit" and follow the UI instructions
4. Send ≥ $5 USDT0 to complete deposit
5. Subaccount created automatically

## Option B: Deposit via SDK

Python

TypeScript

Rust

**Wait time:** Deposits take ~5-10 seconds on Ink Sepolia testnet. Query `subaccount_info` to check when your subaccount is created.

---

## 📈 Step 4: Place Your First Order (1 minute)

Now let's place a simple limit order on BTC-PERP:

Python

TypeScript

**Congratulations!** 🎉 You just placed your first order on Nado! The order is a limit buy 1% below market, so it likely won't fill immediately (it'll rest on the order book).

---

## ✅ Step 5: Verify Your Order

Check that your order was accepted and is on the order book:

Python

TypeScript

Rust

---

## 🎯 What You Just Did

In just 5 minutes, you:

1. ✅ Installed the Nado SDK
2. ✅ Connected to Nado testnet (Ink Sepolia)
3. ✅ Created a subaccount (via deposit)
4. ✅ Placed a limit order on BTC-PERP
5. ✅ Verified the order is on the book

---

## 🚀 Next Steps

## Learn More

* [**Core Concepts**](/developer-resources/get-started/core-concepts) - Deep dive into authentication, subaccounts, linked signers, and signing
* [**First Deposit**](/developer-resources/get-started/first-deposit) - Detailed deposit guide with multiple methods
* [**Linked Signers Guide**](/developer-resources/get-started/linked-signers) - Set up 1-Click Trading and secure bot keys

## Keep Building

* [**First Trade**](https://github.com/nadohq/nado-docs/blob/main/docs/developer-resources/get-started/first-trade.md) - Advanced order types and trading strategies
* [**API Reference**](/developer-resources/api) - Complete API documentation
* [**Common Errors**](/developer-resources/get-started/common-errors) - Troubleshoot signature errors and deposit issues

## SDK Documentation

* [**Python SDK**](https://nadohq.github.io/nado-python-sdk/index.html) - Full Python SDK docs
* [**TypeScript SDK**](/developer-resources/typescript-sdk) - Complete TypeScript SDK guide
* [**Rust SDK**](https://crates.io/crates/nado-sdk) - Rust SDK on crates.io

---

## 💡 Pro Tips

## Cancel Your Test Order

If you want to cancel the order you just placed:

Python

TypeScript

---

## Place a Market Order (Fills Immediately)

Want instant execution? Use a market order by setting an aggressive price:

Python

TypeScript

---

## 🆘 Troubleshooting

## "Signature does not match" Error

See [Common Errors: Error 2028](/developer-resources/get-started/common-errors#error-2028-signature-does-not-match-with-senders-or-linked-signers) for detailed troubleshooting.

**Quick fix:** Verify you're using the correct private key and connected to testnet (Ink Sepolia).

---

## "Subaccount does not exist" Error

See [Common Errors: Error 2024](/developer-resources/get-started/common-errors#error-2024-provided-address-has-no-previous-deposits) for detailed troubleshooting.

**Quick fix:** Make sure you deposited ≥ $5 USDT0 and waited for confirmation.

---

## Need Help?

* **Telegram Community**: [Join our Telegram](https://t.me/+whsZJKpiiVwwNjQ0)
* **Submit a ticket**: [Nado Support](https://nado-90114.zendesk.com/hc/en-us/requests/new?ticket_form_id=52275013155481)

---

**You're ready to trade!** 🎉

This quickstart got you up and running. For production use, read the [Core Concepts](/developer-resources/get-started/core-concepts) and [Linked Signers](/developer-resources/get-started/linked-signers) guides to understand best practices and security considerations.


Last updated 12 days ago

---


# 📘TypeScript SDK

Source: https://docs.nado.xyz/developer-resources/typescript-sdk

The TypeScript SDK contains the following packages:

* `@nadohq/client`: top level client package for the majority of use cases.
* `@nadohq/engine-client`: interactions with the offchain sequencer.
* `@nadohq/indexer-client`: query historical data. e.g: candlesticks, events, etc.
* `@nadohq/trigger-client`: interaction with our trigger service (tp/sl/twap orders).
* `@nadohq/shared`: utility library; e.g. numbers, dates.


Last updated 1 month ago

---


# Getting Started

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/getting-started

## Installation

## Prerequisites

* [Node](https://nodejs.org/en/download/)
* [Yarn](https://classic.yarnpkg.com/lang/en/docs/install/#mac-stable), if using the `yarn` dependency manager
* [Viem](https://viem.sh/)
* [Bignumber.js](https://github.com/MikeMcl/bignumber.js/)

Version 1.x.x of the SDK now uses `viem` instead of `ethers` to represent wallets and RPC connections.

## Install the packages

The Nado SDK packages are hosted on NPM.

## Run the following:

yarn

npm

Copy

```
yarn add @nadohq/client viem bignumber.js
```

Copy

```
npm install @nadohq/client viem bignumber.js
```


Last updated 1 month ago

---


# How To

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to

[Create a Nado client](/developer-resources/typescript-sdk/how-to/create-a-nado-client)[Useful Common Functions](/developer-resources/typescript-sdk/how-to/useful-common-functions)[Query Markets & Products](/developer-resources/typescript-sdk/how-to/query-markets-and-products)[Deposit Funds](/developer-resources/typescript-sdk/how-to/deposit-funds)

Last updated 1 month ago

---


# Create a Nado client

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to/create-a-nado-client

## The `NadoClient` Object

To start using the SDK, you need an initialized `NadoClient` from the `client` package. The `NadoClient` is the main entrypoint to common APIs.

## Create a `NadoClient` object

The `NadoClient` class is rarely instantiated directly. Instead, call the `createNadoClient` function from the `client` package and provide the relevant parameters.

## Import the dependencies

Copy

```
import { createNadoClient } from '@nado-protocol/client';
import { createPublicClient, createWalletClient, http } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import { inkSepolia } from 'viem/chains';
```

## Create a `WalletClient` and `PublicClient`

The `WalletClient` is optional and required only for write operations

Copy

```
const walletClient = createWalletClient({
  account: privateKeyToAccount('0x...'),
  chain: inkSepolia,
  transport: http(),
});

const publicClient = createPublicClient({
  chain: inkSepolia,
  transport: http(),
});
```

## Call `createNadoClient`

The first argument is the `ChainEnv`associated with the client. Each client can talk to one chain that Nado is deployed on. For example, use `inkTestnet`to connect to Nado's instance on Ink Sepolia.

## Full example

Run the script, this example uses `ts-node`:

If no errors are thrown, you're good to go!


Last updated 1 month ago

---


# Deposit Funds

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to/deposit-funds

## Import the functions

We'll use a few of the [common functions](/developer-resources/typescript-sdk/how-to/useful-common-functions), assuming that they are in a `common.ts` file. The withdraw step requires a nonce as the transaction is executed against the off-chain engine.

Copy

```
import { toFixedPoint } from '@nado-protocol/shared';
// Change the import source as needed
import { getNadoClient, prettyPrintJson } from './common';
```

## Mint a mock ERC20 token for testing

Grab a client object and mint mock tokens for the relevant product. This is *only* available on testnets for obvious reasons.

Minting is on-chain, so we wait for the transaction confirmation for chain state to propagate.

Copy

```
const nadoClient = await getNadoClient();
const { walletClient, publicClient } = nadoClient.context

// If you have access to `walletClient`, you can call `walletClient.account.address`
// directly instead of reaching into `nadoClient.context`
const address = walletClient!.account.address;
const subaccountName = 'default';
// 10 USDT0 (6 decimals)
const depositAmount = toFixedPoint(10, 6);

const mintTxHash = await nadoClient.spot._mintMockERC20({
  amount: depositAmount,
  productId: 0,
});

await publicClient.waitForTransactionReceipt({
  hash: mintTxHash,
});
```

## Make a deposit

First, call `approveAllowance` to approve the deposit amount.

This is also an on-chain transaction with a confirmation hash.

Now we can deposit the tokens. This transaction is on-chain.

**Subaccounts**

* A subaccount is an *independent* trading account within Nado, allowing traders to manage risk across independent subaccounts
* Subaccounts are associated by a string `name` (max 12 char.) and the owner wallet address

After this, we inject a short delay while the offchain sequencer picks up the transaction and credits the account.

## Query Subaccount balance

Now, call the `getSubaccountEngineSummary` function to retrieve an overview of your subaccount, including balances.

You should see that your balance associated with `productId` of `0` now reflects your deposit amount.

## Full example


Last updated 1 month ago

---


# Manage Orders

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to/manage-orders

This guide shows you how to:

* Create an order.
* Place an order.
* Cancel an order.
* Query orders by subaccount.

If you inspect the underlying types for these transactions, you'll notice that a `nonce` field is required. This is a unique integer in ascending order. Our off-chain engine has a `nonce` query to return the latest nonce for a given subaccount. All this is abstracted away within the SDK, so you do not need to manually use this query.

## Import the functions

Copy

```
import { getNadoClient, prettyPrintJson } from "./common";
import { nowInSeconds, toFixedPoint, packOrderAppendix } from "@nado-protocol/shared";
```

## Scaffold Your Subaccount

To place orders, we need a subaccount with funds. We need to perform the [deposit funds](/developer-resources/typescript-sdk/how-to/deposit-funds) step as before, this time with 1000 USDT0.

Copy

```
const nadoClient = getNadoClient();
const { walletClient, publicClient } = nadoClient.context;

const address = walletClient!.account.address;
const subaccountName = 'default';
const depositAmount = toFixedPoint(1000, 6);

const mintTxHash = await nadoClient.spot._mintMockERC20({
  amount: depositAmount,
  productId: 0,
});
await publicClient.waitForTransactionReceipt({
  hash: mintTxHash,
});

const approveTxHash = await nadoClient.spot.approveAllowance({
  amount: depositAmount,
  productId: 0,
});
await publicClient.waitForTransactionReceipt({
  hash: approveTxHash,
});

const depositTxHash = await nadoClient.spot.deposit({
  subaccountName: 'default',
  amount: depositAmount,
  productId: 0,
});
await publicClient.waitForTransactionReceipt({
  hash: depositTxHash,
});

await new Promise((resolve) => setTimeout(resolve, 10000));
```

## Create an order

Placing an order requires a number of parameters, represented by the `PlaceOrderParam['order']` type.

In the example below:

* The order `appendix` indicates order execution type and other flags. Please refer to [Order Appendix](/developer-resources/api/order-appendix) for more details.
* The order `expiration` time is given by calling the `nowInSeconds` function from the `utils` package and adding 60 seconds. This means the order will expire 60 seconds from now.
* The `price` field is set at `80000` - a low value (at the time of writing) to prevent execution. This enables us to cancel the order later on without it being instantly filled. Please adjust this price accordingly.
* The `amount` field is set at `10**16` - this is the amount to buy/sell. A positive value is to buy, negative is to sell.

  + Amount is normalized to 18 decimal places, which is what `toFixedPoint` does by default.
  + **NOTE**: Min limit order size for `BTC` is `10**16` and for `ETH` is `10**17`. Orders below these sizes will fail to be placed.

## Place the order

Use the order parameters to place the order with the `placeOrder` function.

## Alternative order placement

Alternatively, you can manually use `payloadBuilder` to manually generate the place order payload. This may be useful in cases where you want to build the `tx` separately from sending the execute API call.

## Order digest

You can optionally generate the order digest, which can then be used to further manage the order e.g: cancelling the order. The order digest is also returned upon executing the `placeOrder` transaction.

## Query orders on the subaccount

Now we can query the subaccount for open orders with the `getOpenSubaccountOrders` function.

## Cancel order

Cancel the order using the digest of the placed order. You can cancel multiple orders at once.

## Query Orders to Verify Cancellation

Run [query orders on the subaccount](/developer-resources/typescript-sdk/how-to/manage-orders#query-orders-on-the-subaccount) again to make sure the cancellation was successful.

## Clean up

Finally, clean up by withdrawing the same amount as you have deposited, minus the 1 USDT0 withdrawal fee.

## Full example

[PreviousWithdraw Funds](/developer-resources/typescript-sdk/how-to/withdraw-funds)

Last updated 1 month ago

---


# Query Markets & Products

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to/query-markets-and-products

In this section, we'll be going over fetching:

* State and config for all markets & products.
* Latest market price for one product.
* Market liquidity for one product (i.e. amount of liquidity at each price tick).

For all available queries, consult the [API reference](https://nadohq.github.io/nado-typescript-sdk/).

## All Markets Query

The `getAllEngineMarkets` function returns the state of all markets from our backend API, which reflects the state of the off-chain matching engine.

Copy

```
// Fetches state from offchain sequencer 
const allMarkets = await nadoClient.market.getAllMarkets();
```

## Latest market price

The `getLatestMarketPrice` function returns the market price data of a single product given by its product id.

Copy

```
const latestMarketPrice = await nadoClient.market.getLatestMarketPrice({
  productId: 1,
});
```

## Market liquidity

The `getMarketLiquidity` function returns the available liquidity at each price tick. The number of price levels for each side of the book is given by `depth`. For example, a depth of `2` will retrieve 2 levels of bids and 2 levels of asks. Price levels are separated by the `priceIncrement` of the market, given by the `getAllMarkets` query.

## Full example


Last updated 1 month ago

---


# Useful Common Functions

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to/useful-common-functions

These are some utility functions used throughout the guide. You may want to include them in your project.

* A `getNadoClient` function that returns a ready-made client object connected to Ink Sepolia.
* A `prettyPrintJson` function that logs readable JSON.

**Note**

* Make sure your account has funds for gas on the relevant network.
* Make sure to replace your private key of choice in the function `getNadoClient` below

Copy

```
import { createNadoClient } from '@nado-protocol/client';
import { toPrintableObject } from '@nado-protocol/shared';
import { createPublicClient, createWalletClient, http } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import { inkSepolia } from 'viem/chains';

/**
 * Creates a Nado client for example scripts
 */
export function getNadoClient() {
  const walletClient = createWalletClient({
    account: privateKeyToAccount('0x...'),
    chain: inkSepolia,
    transport: http(),
  });

  const publicClient = createPublicClient({
    chain: inkSepolia,
    transport: http(),
  });

  return createNadoClient('inkTestnet', {
    walletClient,
    publicClient,
  });
}

/**
 * Util for pretty printing JSON
 */
export function prettyPrintJson(label: string, json: unknown) {
  console.log(label);
  console.log(JSON.stringify(toPrintableObject(json), null, 2));
}
```


Last updated 1 month ago

---


# Withdraw Funds

Source: https://docs.nado.xyz/developer-resources/typescript-sdk/how-to/withdraw-funds

This section assumes you have followed the instructions and run the code in the [deposit funds example](/developer-resources/typescript-sdk/how-to/deposit-funds).

## Call the `withdraw` function

Note that there is a fee for withdrawals, in our case, it should be 1 USDT0. The `amount` sent to the `withdraw` function is exclusive of fee, hence depositAmount - toFixedPoint(1, 6)

Copy

```
const withdrawTx = await nadoClient.spot.withdraw({
  amount: depositAmount - toFixedPoint(1, 6),
  productId: 0,
  subaccountName: 'default',
});

prettyPrintJson('Withdraw Tx', withdrawTx);
```

## Check balances

Retrieve and log the subaccount balances using `getSubaccountSummary`. You should see that your balance for USDT0 (product ID of 0) is now 0.

## Full example - deposit and withdraw

Copy

```
import { toFixedPoint } from '@nado-protocol/utils';
import { getNadoClient, prettyPrintJson } from './common';

async function main() {
  const nadoClient = getNadoClient();
  const { walletClient, publicClient } = nadoClient.context;

  // If you have access to `walletClient`, you can call `walletClient.account.address`
  // directly instead of reaching into `nadoClient.context`
  const address = walletClient!.account.address;
  const subaccountName = 'default';
  // 10 USDT0 (6 decimals)
  const depositAmount = toFixedPoint(10, 6);

  // TESTNET ONLY - Mint yourself some tokens
  const mintTxHash = await nadoClient.spot._mintMockERC20({
    amount: depositAmount,
    productId: 0,
  });
  // Mint goes on-chain, so wait for confirmation
  await publicClient.waitForTransactionReceipt({
    hash: mintTxHash,
  });

  // Deposits require approval on the ERC20 token, this is on-chain as well
  const approveTxHash = await nadoClient.spot.approveAllowance({
    amount: depositAmount,
    productId: 0,
  });

  await publicClient.waitForTransactionReceipt({
    hash: approveTxHash,
  });

  // Now execute the deposit, which goes on-chain
  const depositTxHash = await nadoClient.spot.deposit({
    // Your choice of name for the subaccount, this subaccount will be credited with the deposit balance
    subaccountName: 'default',
    amount: depositAmount,
    productId: 0,
  });

  await publicClient.waitForTransactionReceipt({
    hash: depositTxHash,
  });
  
  await new Promise((resolve) => setTimeout(resolve, 10000));

  const subaccountData =
    await nadoClient.subaccount.getSubaccountSummary({
      subaccountOwner: address,
      subaccountName,
    });
  prettyPrintJson('Subaccount Data After Deposit', subaccountData);

  // Now withdraw your funds, this goes to the off-chain engine
  // We're withdrawing less than 10 as there are withdrawal fees of 1 USDT0
  const withdrawTx = await nadoClient.spot.withdraw({
    amount: depositAmount - toFixedPoint(1, 6),
    productId: 0,
    subaccountName,
  });
  prettyPrintJson('Withdraw Tx', withdrawTx);

  // Your new subaccount summary should have zero balances
  const subaccountDataAfterWithdraw =
    await nadoClient.subaccount.getSubaccountSummary({
      subaccountOwner: address,
      subaccountName,
    });
  prettyPrintJson(
    'Subaccount Data After Withdrawal',
    subaccountDataAfterWithdraw,
  );
}

main();
```


Last updated 1 month ago

---


# FAQs

Source: https://docs.nado.xyz/faqs

Welcome to the Nado Frequently Asked Questions (FAQs). This section provides detailed, practical answers for newcomers to Nado. Whether you're new to perpetual trading or refining your strategy, these responses aim to clarify key concepts and processes.

If you're depositing into Nado or withdrawing for the first time, we recommend reviewing the relevant responses and associated tutorial sections for a smooth onboarding experience.

Trade smarter, move faster.

---

## Deposits

## Why can't I deposit my ETH?

Nado exclusively accepts ERC-20 tokens as deposits to ensure seamless integration with our smart contract infrastructure on the INK network.

Native ETH is not directly supported for deposits into trading positions, as it serves primarily as the gas token for transactions on INK. To use ETH-based value for trading, first bridge or wrap it into a compatible ERC-20 asset (e.g., Wrapped ETH or stablecoins like USDT0) on the Ink network. This process can be initiated through trusted bridges accessible via the Nado interface or external tools, allowing you to convert and deposit in an efficient workflow.

Users can also send ETH to Nado via the direct-deposit flow available on the app, simplifying the deposit process.

---

## How do I deposit from other chains onto Ink and deposit on Nado?

Nado operates on the high-performance Ink L2 network, so deposits from other chains require first bridging assets to Ink before depositing them into Nado. This two-step process ensures compatibility and leverages Ink’s low-cost, EVM-compatible environment.

1. **Bridge Assets to INK**: Use a reliable cross-chain bridge to transfer supported assets (e.g., ETH, USDT0, USDC) from your source chain (such as Ethereum, Arbitrum, Polygon, etc) to your Ink wallet address.

   1. Recommended Bridges: [USDT0 Native Bridge](https://usdt0.to/transfer), [Superbridge](https://superbridge.app/), [Bungee](https://bungee.exchange/), or [Relay](https://relay.link/).
   2. Steps:

      1. Connect your wallet and select the source chain.
      2. Choose Ink as the destination network.
      3. Select the asset and amount, then sign the bridging transaction.
      4. Wait for confirmation (typically 5 - 15 minutes); track via the [Ink Explorer](https://explorer.inkonchain.com).

Ensure your wallet is configured for INK.

1. **Deposit into Nado**: Once assets are on Ink and visible in your wallet:

   1. Visit the [Nado app](https://app.nado.xyz).
   2. Connect your Ink-configured wallet.
   3. Navigate to the Deposit section on the Portfolio page.
   4. Select the asset, enter the amount, and confirm the on-chain transaction – deposits settle near-instantly on Nado, making collateral available for trading almost immediately.

This method supports assets like ETH (for gas), USDT0, USDC, wETH, and wBTC. Always verify token contract addresses on the Ink Explorer to avoid errors.

---

## How do I deposit funds from a CEX wallet?

Depositing from a centralized exchange (CEX) to Nado involves withdrawing assets to your Ink-configured wallet first, then depositing into Nado This ensures funds land on the correct network for seamless trading.

**1. Withdraw from CEX to Ink Wallet**:

* Log in to your CEX account (e.g., Kraken) and initiate a withdrawal of supported assets (e.g., ETH, USDT0, etc).
* Enter your Ink wallet address as the recipient – double-check it's on the Ink network to prevent loss of funds. You can also use the 'Direct Deposit' feature on Nado to simplify this process.
* Specify the Ink network in the withdrawal options (Kraken supports zero-fee Ink withdrawals for ETH).
* Confirm and complete the withdrawal; funds typically arrive in 10–30 minutes.

**2. Import and Verify Assets in Wallet:**

* If the token doesn't auto-appear, manually import it using the INK-specific contract address (e.g., via MetaMask's "Import Token" feature—verify addresses on the Ink Explorer).

**3. Deposit into Nado:**

* Connect your wallet to the [Nado app](https://app.nado.xyz) on the Ink network.
* Go to the Deposit section in the Portfolio.
* Select the asset, input the amount, and approve the transaction – your collateral will be available almost immediately for trading.

**Pro Tip**: Start with a small test withdrawal to confirm the process. For gas fees on Ink, ensure you have at least 0.01 ETH in your wallet.

---

## Withdrawals

Withdrawals from Nado are designed for simplicity and efficiency, operating without the need for any bridging mechanisms. This means your funds move directly from the protocol to your connected wallet in a straightforward, on-chain process, preserving the integrity and speed of your assets' transfer.

To enhance your withdrawal flexibility, you have the option to enable the borrowing toggle. This feature allows you to borrow additional assets against your existing margin, providing greater liquidity while you initiate the withdrawal.

> Once your withdrawal request is submitted, it joins a batch of other pending withdrawals for optimized processing. You can easily monitor the real-time status of your withdrawal – including confirmation, pending settlement, and completion – directly in the **Withdrawals History** tab, accessible via the Portfolio page in the Nado interface.

Nado employs a gas-optimization strategy to minimize user fees, batching and submitting withdrawal transactions to the Ink L2 network only when gas prices are at their lowest.

While all user actions within Nado execute instantaneously, the actual on-chain settlement of withdrawals may take up to 30 minutes during periods of elevated network congestion. This 30-minute window serves as the targeted maximum pending time: Nado's automated system will proactively submit the transaction to Nado after this interval, even in high-gas environments, to ensure timely resolution.

In rare cases where processing exceeds this timeframe, it is typically attributable to sustained spikes in network gas costs beyond Nado's control. Rest assured, if your withdrawal appears as "pending" in the Nado app, it has been successfully queued and validated internally. Settlement will occur automatically once gas fees decrease or fall below the predefined optimization threshold at the time of submission.

Please be aware that exact withdrawal times can vary due to real-time fluctuations in network conditions and gas pricing. For comprehensive tracking, we also recommend reviewing the Account History section within the Portfolio page, which provides a detailed log of all deposit and withdrawal activities.

---

## General FAQs

## Why do I need to deposit?

Depositing collateral into Nado's smart contracts is essential for enabling leveraged trading on the exchange. These contracts operate in a non-custodial manner, meaning your funds remain under your control and are stored on-chain. You can withdraw your available balance whenever you choose, providing full flexibility while powering your positions.

---

## Do I control my assets?

Yes, you maintain complete control over your assets in Nado. Only you can initiate trades, access funds, or execute withdrawals. The protocol's non-custodial smart contracts ensure that no third party holds custody, giving you sovereign authority over your portfolio.

---

## What is unified margin trading and how is it unique?

Unified margin consolidates all your account balances and open positions into one margin pool, enabling real-time margin offsets that reduce overall margin requirements and enhance capital efficiency.

**How it Works**

* **Shared Collateral Pool**: Every asset in your Nado account (e.g., USDT0 deposits, kBTC holdings, etc) and all open positions contribute to a unified health score, calculated by Nado's on-chain risk engine. This score reflects your total available collateral and maintenance margin levels before hitting margin limits and liquidation thresholds.
* **Automatic Netting & Rebalancing**: Positive PnL from one position can offset losses in another, dynamically adjusting margin needs. Maintenance margin (the buffer against liquidation) and initial margin (required to open positions) are computed holistically, often resulting in lower margin requirements than isolated margin trades.
* **Risk Tiers for Visibility**: Monitor your portfolio through intuitive risk levels, updated in real-time for proactive management.

> The core advantage of unified margin is maximizing capital efficiency.

Your portfolio operates as an interconnected system, where balanced exposures minimize tied-up funds and amplify buying power. It's suited for strategies like basis trades or multi-asset hedges, eliminating the need for constant rebalancing.

---

## How do I earn interest on deposits?

Interest is earned automatically on all asset deposits into Nado, as they are integrated into the protocol's underlying money markets. These markets facilitate leveraged spot trading and borrowing opportunities, with your deposits participating passively to generate yield.

The on-chain smart contracts ensure that borrowers always meet margin requirements, maintaining system stability and security for all participants.

---

## What are the fees?

The fee model follows a classic maker-taker structure across spot and perpetuals markets:

* **Makers** (limit orders adding liquidity): Earn rebates at higher tiers.
* **Takers** (market orders removing liquidity): Pay a modest fee.

All fees are calculated in basis points (bps, or 0.01%) of the trade's notional value and settled instantly in USDT0 from your collateral.

Setting Nado apart from fixed-rate models, the volume-based scaling tiers update monthly to encourage deeper orderbooks and sustained activity. As your 30-day trading volume (maker + taker) climbs, taker fees decrease and maker rebates increase, creating a virtuous cycle of growing efficiency and participation.

> For complete fee details including minimum fees for small taker orders, see the [Fees & Rebates](/fees-and-rebates) page.

---

## Are there take-profit and stop-loss orders?

Yes, Nado supports take-profit (TP) and stop-loss (SL) orders for open perpetual positions. These conditional order types allow you to automatically exit positions at predefined price levels, helping manage risk and secure profits in volatile markets.

Nado also offers traders more advanced order types, including TWAP orders and scale orders (coming soon).

---

## Why do I have less funds available than what I deposited? (“I deposited $50, why does it say I only have $40 to trade with?”)

In Nado's unified margin engine, the Available Margin metric represents the value of your deposited collateral, adjusted by the initial margin weights of each asset. This weighting accounts for varying levels of volatility across collateral types, applying a discount to more volatile assets to ensure prudent risk management. As a result, not all collateral contributes at full face value to your trading capacity.

> For example, stable assets like USDT0 are weighted 1:1 with their nominal value, providing direct usability. This system allows you to leverage diverse collateral types while maintaining overall account stability.

---

## Is there a minimum amount to trade?

Nado imposes a $5 equivalent initial deposit minimum for a user's trading account. After the initial deposit, there is no universal minimum trade amount.

> However, certain order types enforce a minimum order size based on each product's parameters. Additionally, all taker orders are subject to a minimum fee calculated based on the product's minimum size, even for order types that allow smaller sizes. See the [Fees & Rebates](/fees-and-rebates) page for complete details on minimum fees and order requirements.

---

## Why is there a negative sign in front of my asset balance?

A negative balance indicator for an asset signifies that you are currently borrowing that asset within the protocol. This can occur as part of position management, such as when leveraging borrows to enter or maintain trades.

---

## I didn’t borrow USDT0. Why is my balance negative?

A negative USDT0 balance can arise automatically in scenarios involving perpetual positions with unrealized negative PnL, especially when using collateral assets that are not USDT0.

Throughout the duration of holding such positions, the protocol settles PnL in USDT0 between winning and losing trades. If sufficient USDT0 is unavailable in your account, Nado will borrow it on your behalf to cover the settlement, resulting in a temporary negative balance.

---

## How do I repay borrows?

To repay any outstanding borrows in Nado, locate and select the Repay button through one of the following access points:

* **Balances Table**: Click the drop-down menu on the right-most side of the relevant row.

You have two primary options for repayment:

1. **Direct Deposit**: Deposit the exact amount of the borrowed asset to settle the balance in full.
2. **Asset Conversion**: Sell or convert another held asset (e.g., swap wETH for USDT0) to generate the necessary funds for repayment.

This process restores your balance to positive and frees up additional margin for trading.

---

## Why was my position liquidated if the chart shows that the price didn’t hit my Liq. price?

The Liq. Price displayed in the Perp Positions table is an estimated value calculated based on your current account state and the specific position's health. In a multi-position portfolio, fluctuations in other open positions can indirectly affect this estimate, causing it to shift without direct price movement in the charted asset.

> **Note**: The trading terminal chart on the app uses the traded price on Nado, but liquidations use the oracle price.

Liquidations are triggered solely by the market's Oracle Price, sourced from the Chaos Labs Oracle and submitted to the on-chain smart contracts at regular time intervals or in response to significant price changes.

If the Oracle Price causes your account to fall below maintenance margin requirements, liquidation occurs regardless of the displayed estimate. This ensures objective, real-time risk management aligned with on-chain data.

---

## How do liquidations work?

When a subaccount's maintenance health falls below zero, Nado initiates liquidation to restore solvency, closing elements in a structured sequence that minimizes market impact. Any user can act as a liquidator by submitting a transaction to purchase discounted assets or cover marked-up liabilities, earning a profit while aiding recovery.

> The process pauses if initial health rises above zero at any step, giving positions a chance to rebound.

Liquidators specify a product and amount to target, with the system rounding down to the optimal size that brings maintenance health back to non-negative, whilst making sure the initial health is non-positive. – balancing efficiency with user protection.

The sequence of liquidation operations is as follows:

1. **Cancel Open Orders**: All pending orders in the subaccount are voided to free up resources.
2. **Liquidate Assets**: Spot balances, long spreads, and positive-PnL perpetuals are sold at a discount.
3. **Liquidate Liabilities**: Borrows and short spreads are repaid at a markup.

---

## What is Maintenance Margin Usage?

Maintenance Margin Usage is an indicator of when liquidation begins if you hit 100%. It indicates the percentage of your maintenance margin that is consumed by open positions. It provides a real-time gauge of how close your account is to liquidation thresholds.

* **Low Risk**: 0 – 40%
* **Medium Risk**: 40 – 70%
* **High Risk**: 70 – 90%
* **Extreme Risk**: 90 – 100%

> If Maintenance Margin Usage reaches 100%, your account is immediately eligible for liquidation. At this point, you will be unable to open new positions until some margin is freed up, either through position closures, additional deposits, or positive PnL realizations.

Monitoring this metric helps prevent overexposure and ensures you retain capacity for opportunistic trades.

---

## What is Available Margin?

Available Margin quantifies the amount of tradable funds or collateral in your account, calculated as the initial weighted margin that remains unused. This represents your immediate buying power for initiating new positions.

Should Available Margin reach $0, new position openings will be restricted. It is also commonly referred to as Free Collateral, serving as a key indicator of your account's liquidity for trading.

---

## What is Maintenance Margin?

Maintenance Margin represents the buffer of funds or collateral in your account before it reaches liquidation eligibility. If this value drops to $0, your account enters a high-risk state and may be subject to automated liquidation.

Maintaining a positive Maintenance Margin is crucial for position sustainability — regularly review this alongside market conditions to adjust leverage proactively.

You must maintain a Maintenance Margin value above $0 to avoid liquidation.

---

## What are Initial and Maintenance weights?

In exchanges limited to dollar-pegged collateral (e.g., stablecoins), assets are typically weighted at full face value for simplicity. However, Nado's cross-margin system accepts multiple collateral types and applies dual weights to account for volatility, ensuring robust risk controls:

1. **Initial Weight**: Determines the collateral value available for opening new positions (i.e., trading capacity).
2. **Maintenance Weight**: Sets the threshold for sustaining positions without triggering liquidation.

These weights provide traders with clear visibility into both offensive (trading) and defensive (risk) aspects of their portfolio.

---

## Initial vs. Maintenance Margin

Initial and maintenance weighted margins offer traders dual insights into account health: your capacity to enter trades and proximity to liquidation risks.

* **Initial Margin**: The total funds available for trading, computed as initial weighted collateral minus initial weighted margin requirements.
* **Maintenance Margin**: The minimum funds required to avoid liquidation, calculated as maintenance weighted collateral minus maintenance weighted margin requirements.

---

## Are there deposit caps on the NLP during the Private Alpha?

Yes, during Private Alpha, **the NLP** ***deposit amount*** **is capped at 20,000 USDT0 per Nado trading account**. This means users cannot *add* more than 20,000 USDT0 into the vault.

However, **the position value itself can grow beyond 20,000 USDT0 through yield**. For example, if a user deposits 19,000 USDT0 and their position appreciates to 20,000 USDT0, that’s allowed, they simply can’t deposit additional funds once they’ve hit the deposit limit.

Any updates to the deposit cap will be communicated in advance and reflected in the Nado app during Private Alpha or later.

---

## What is the deposit APY for the NLP?

Yields for LPs in the NLP vault are variable and subject to change based on the prevailing market conditions and other factors including but not limited to the vault strategy’s PnL and the total LP capital (USDT0) deposited into the vault.

---

For any other questions or feedback, please refer to the Nado community and support channels for assistance.


Last updated 17 days ago

---


# Fees & Rebates

Source: https://docs.nado.xyz/fees-and-rebates

Nado's fee structure is designed for precision and fairness, rewarding liquidity providers while keeping costs lean for active traders.

The fee model follows a classic maker-taker structure across spot and perpetuals markets:

* **Makers** (limit orders adding liquidity): Earn rebates at higher tiers.
* **Takers** (market orders removing liquidity): Pay a modest fee.

All fees are calculated in basis points (bps, or 0.01%) of the trade's notional value and settled instantly in USDT0 from your collateral.

Setting Nado apart from fixed-rate models, the volume-based scaling tiers update monthly to encourage deeper orderbooks and sustained activity. As your 30-day trading volume (maker + taker) climbs, taker fees decrease and maker rebates increase, creating a virtuous cycle of growing efficiency and participation.

Traders can monitor their current tier and monthly volume accrual directly in the account overview, with epochs resetting on the first of each UTC month.

Fee tiers compete with DeFi's most cost-effective venues, letting high-volume users pay as little as 1.5 bps as a taker – while elite makers are compensated to bolster the book.

---

## Trading Fees & Scaling Tiers

Nado's fee tiers scale based on your total trading volume (maker + taker) over the prior 30 days, aggregated across spot and perpetuals. Starting at accessible entry levels, they progress to elite rebates, encouraging consistent engagement without arbitrary barriers. Benefits include:

* **Capital Efficiency**: Lower fees free up more funds for positioning, amplifying returns on scale.
* **Liquidity Incentives**: Maker rebates (expressed as negative fees) directly reward orderbook depth, tightening spreads for all traders.
* **Transparency**: Real-time tier visibility lets you strategize volume to unlock better rates, with no hidden multipliers.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-45c07a137e9ec7ef54c0d4c0def7e1b66c6e6f5b%252FFee%2520%2526%2520Rebates%2520Tables02.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=d63713b6&sv=2)

Fees apply per execution:

* For a partial fill, only the matched portion incurs fee charges.
* No fees are applied to deposits or non-trading actions like borrowing.

## Minimum Fee for Takers

Taker orders are subject to a minimum fee based on each product's minimum size:

* **Orders ≥ minimum size**: Standard fee applies to the full notional value: `orderSize × feeRate`
* **Orders < minimum size**: Minimum fee applies: `minSize × feeRate` (effectively treating small orders as if they were minSize)

The minimum size is measured in **notional dollar value** (USDT0), not the asset quantity itself. For example, if minSize is $100, you could trade 0.001 BTC at $100,000/BTC (= $100 notional) or 0.033 ETH at $3,000/ETH (= $100 notional).

While certain order types enforce a minimum order size, the minimum fee applies to ALL taker orders regardless of size. This ensures consistent fee collection even for order types that allow smaller sizes.

**Example - Large order**: Assume minSize = $100. A $50,000 taker order at 3.5 bps:

* Fee: $50,000 × 0.035% = $17.50 (standard calculation)

**Example - Small order**: Assume minSize = $100. A $75 taker order at 3.5 bps:

* Fee: $100 × 0.035% = $0.035 (minimum fee applies, calculated as if the order was minSize)

Maker orders do not have minimum fee requirements and are charged or receive rebates on the full notional value based on your fee tier (see fee table above).

## Example: Impact of Fee Tiers on a Trade

Consider a $50,000 ETH perpetuals market order (taker) or limit order (maker) at ETH $3,000.

> **Entry Tier** ($0 volume): Taker pays 3.5 bps = $17.50 (0.035% × $50,000). Maker pays 1 bp = $5.00.

> **Mid-Tier** ($25M volume): Taker pays 3 bps = $15.00. Maker pays 0.5 bps = $2.50.

> **Elite Tier** ($5B+ volume): Taker pays 1.5 bps = $7.50. Maker receives rebate of 0.8 bps = -$4.00 (added to collateral).

The scaling dynamic rewards whales without alienating newcomers, as even base rates (3.5 bps taker) beat many other exchange venues fee models.

Maker rebates accrue instantly to your subaccount, boosting health and compounding on unified cross-margin.

---

## Sequencer & Network Fees

Nado's off-chain sequencer handles order matching for sub-15 ms speed, with on-chain settlement via Ink L2.

To cover onchain interactions without gas volatility, Nado charges flat, predictable fees in USDT0 (or equivalent asset for withdrawals) – cheaper than typical L2 costs.

* **Deposits**: 0 USDT0 (instant bridging from Ink).
* **Order Placement** (Maker): 0 USDT0.
* **Order Placement** (Taker): 0 USDT0 (bundled with trading fee).
* **Subaccount Transfers**:

  + Standard: 1 USDT0
  + Isolated subaccounts: 0.1 USDT0 (when either sender or recipient is isolated)
* **Liquidations**: 1 USDT0 (to the liquidator).
* **Withdrawals**:

  + USDT0: 1 USDT0
  + wETH: 0.0006 wETH
  + kBTC: 0.00004 kBTC

Sequencer fees are exclusively applied upon successful execution of the corresponding action. Failed actions, such as an attempt at an under-collateralized withdrawal, incur no sequencer fee charge.

Nado submits orders on-chain to the Ink L2 in batches to optimize for gas efficiency. During instances of excessive network congestion and high gas costs on-chain, Nado withdrawals may queue longer than usual, typically settling within ~30 minutes maximum.

## Fast Withdrawals

Users can also select the option for "Fast Withdrawals" anytime for withdrawal priority, which incurs an extra USDT0 charge, calculated as follows:

Fast Withdrawal Fee=1 USDT0+max⁡(5 USDT0,10 bps×amount)\text{Fast Withdrawal Fee} = 1 \, \text{USDT0} + \max\left(5 \, \text{USDT0}, 10 \, \text{bps} \times \text{amount}\right)Fast Withdrawal Fee=1USDT0+max(5USDT0,10bps×amount)

For BTC, the fee calculation is:

Fast Withdrawal Fee (BTC)=1 USDT0+max⁡(0.0002 BTC,10 bps×amount), when the withdrawal fee for BTC is 0.00004 BTC\text{Fast Withdrawal Fee (BTC)} = 1 \, \text{USDT0} + \max\left(0.0002 \, \text{BTC}, 10 \, \text{bps} \times \text{amount}\right), \text{ when the withdrawal fee for BTC is } 0.00004 \, \text{BTC}Fast Withdrawal Fee (BTC)=1USDT0+max(0.0002BTC,10bps×amount), when the withdrawal fee for BTC is 0.00004BTC

With tiered scaling and minimal overhead, Nado's fee model turns every fill into forward momentum – precise, scalable, and built for the long-haul.

---


Last updated 23 days ago

---


# Funding Rates

Source: https://docs.nado.xyz/funding-rates

Funding rates are a core feature of perpetual contracts on Nado, designed to keep the price of the contracts closely aligned with the underlying spot price of the asset, even though perpetuals have no expiration date.

Unlike traditional futures, which naturally converge to spot prices as expiry approaches, perpetuals rely on funding rates to achieve this balance. These rates represent small, periodic payments exchanged between long and short positions – longs pay shorts when the perpetual price exceeds the spot price (positive rate), and shorts pay longs in the opposite case (negative rate).

> The funding rate mechanism encourages trading activity that pulls prices back toward equilibrium, ensuring fair and efficient markets.

On Nado, funding rates settle every hour and are proportional to your position's notional value – meaning the full size of the trade, not just your margin. Rates are displayed in real-time on the trading terminal and are capped at 2% per day to avoid excessive swings.

By incorporating funding rates into your strategy, you can anticipate holding costs or potential rebates on open positions, turning what might seem like a minor detail into a key factor for long-term trades.

---

## How Funding Rates Work

Nado calculates funding rates using a transparent formula that compares the perpetual contract's mark price to the spot index price, sourced from third-party oracles. This difference, known as the funding index, determines the rate paid or received at each funding interval.

The goal is to mimic the natural price convergence of expiring contracts, but continuously, fostering balanced liquidity without fixed deadlines.

## Key Components

* **Spot Index Price**: An aggregated benchmark from major exchanges (Binance, Coinbase Pro, Kraken) via the Stork oracle network. It uses a median of real-time prices to resist manipulation, providing a stable reference for the asset's true market value.
* **Perpetual Mark Price**: A time-weighted average price (TWAP) calculated from Nado's orderbook over the funding interval. This reflects on-chain trading activity without relying solely on external oracles, ensuring accuracy tied to actual liquidity.
* **Funding Interval**: Payments occur hourly, on the hour, between all open positions in the market. This frequent adjustment keeps deviations minimal.
* **Calculation Formulas**:

> **Payment Mechanics**: The rate applies to the notional value of your position. For example, a 0.01% hourly rate on a $10,000 notional long means paying (or receiving) $1 that hour, impacting a user’s unsettled USDT0. No platform fees are involved – it's a direct transfer between longs and shorts.

These components work together like a thermostat in a room: the index detects temperature (price) deviations from the set point (spot), and the rate adjusts the "heat" (incentives) to restore balance, operating smoothly in the background.

---

## Positive Funding Rates

A positive funding rate occurs when the perpetual mark price trades above the spot index price, indicating stronger demand for longs. In this scenario, holders of long positions pay shorts, which discourages excessive bullishness and encourages shorts to enter, gradually pushing the perpetual price down toward spot levels.

> This dynamic is common in bullish markets, where optimism drives perps higher. The funding rate acts as a counterbalance, rendering extended long positions more expensive while rewarding shorts for providing liquidity.

**Example**: Suppose ETH's spot index price is $3,000, but Nado's ETH perpetual mark price is $3,030 (a 1% premium). This yields a positive funding rate of +0.01% per hour (or about 3.65% annualized).

* Alice holds a long position of 10 ETH perpetuals, with a notional value of $30,300 (10 × $3,030).
* Alice’s initial margin is $3,030 (10x leverage).
* At settlement, Alice pays shorts: 0.01% of $30,300 = $3.03 for the hour.

Over 24 hours, this totals about $72.72, subtracted from Alice’s collateral.

If ETH rises 2% during the day (to $3,060 spot), Alice’s position gains $606 in PnL. Subtracting the $72.72 funding cost, and Alice’s net gain is $533.28 – a meaningful reduction in PnL for holding the position over extended periods.

In contrast, a short position of the same size would receive that $72.72 funding rate payment, boosting their returns as the premium unwinds. Historically, such rates can persist for weeks during a market uptrend, so monitoring the funding rates helps decide whether to close, hedge, or flip to short for the rebate.

---

## Negative Funding Rates

When the perpetual mark price falls below the spot index price, the funding rate turns negative, meaning shorts pay longs. This setup counters bearish pressure by rendering short positions more costly to keep open, attracting longs to buy the "discount" and lift prices back up.

> Negative rates often emerge in market downtrends or when spot sentiment outpaces perpetuals, providing a subtle boost to longs and a reminder for shorts to reassess.

**Example**: Now, ETH's spot index is $3,000, but the perpetual mark is $2,970 (a 1% discount), resulting in a -0.01% hourly rate.

* Alice’s Long Position: 10 ETH perpetuals, notional $29,700 (10 × $2,970).
* Initial Margin: $2,970 (10x leverage).
* Shorts Pay Alice: 0.01% of $29,700 = $2.97 for the hour.
* Daily Total: about $71.28 added to Alice’s collateral.

If ETH drops 2% (to $2,940 spot), Alice’s position loses $594 in PnL. The $71.28 funding receipt softens this to a net loss of $522.72. For a short of the same size, they'd pay out $71.28, compounding their costs in a falling market.

In extended bear phases – where negatives have lasted months – this can turn defensive longs into earners, offsetting volatility and rewarding patience.

Funding rates evolve with open interest and market sentiment, so make sure to review them in Nado's trading terminal before entering trades. By factoring in funding rates, traders can align their strategies with these built-in incentives, enhancing precision across spot, perps, and beyond.

---


Last updated 1 month ago

---


# Legal

Source: https://docs.nado.xyz/legal

[Restricted Territories](/legal/restricted-territories)


Last updated 1 month ago

---


# Restricted Territories

Source: https://docs.nado.xyz/legal/restricted-territories

Nado is not available in certain jurisdictions. Access to the platform is determined by the location from which a person connects. Depending on the region, persons may be fully blocked from accessing Nado or may be able to view the interface but not interact with it. If you attempt to access or interact with Nado from a restricted region, the interface will limit or block functionality accordingly.

## Fully Restricted Territories

Persons located in the following jurisdictions cannot access or use Nado:

* Belarus
* Cuba
* Iran
* North Korea
* Russia
* Ukraine

## View-Only Territories

Persons located in the following jurisdictions may access and view the interface only. Trading and all interactive features are disabled:

* Canada
* United States of America

Please refer to the Terms of Use for further details.


Last updated 27 days ago

---


# Liquidations

Source: https://docs.nado.xyz/liquidations

Liquidations on Nado serve as a critical safety valve in leveraged trading, automatically closing positions when your subaccount's maintenance health dips below zero – ensuring losses never exceed your deposited collateral and preserving the platform's stability for all users.

> The liquidation process, triggered by on-chain calculations, uses the mark oracle price from the Chaos Oracle network for fairness, drawing from a time-weighted average of third-party exchange data.

Nado's design emphasizes precision where liquidations prioritize minimal disruption, starting with the most liquid assets and halting if health recovers mid-process. Complementing the liquidation mechanism is the insurance fund, a dedicated USDT0 reserve that absorbs rare shortfalls, acting as the platform's first line of defense.

The insurance fund is sustained by 50% of all liquidation profits. It helps to prevent widespread loss socialization, maintaining platform solvency even during extreme volatility.

---

## How Liquidations Work

When a subaccount's maintenance health falls below zero, the account becomes *eligible* for liquidation. Nado does **not** automatically liquidate it; a liquidator must submit a transaction to kick off the flow. Any user can do this by calling the [`liquidate-subaccount` API endpoint](/developer-resources/api/gateway/executes/liquidate-subaccount) or the on-chain contract method to buy discounted assets or repay marked-up liabilities, earning the spread while restoring solvency. Incentives keep liquidators active, so liquidations are typically picked up quickly.

> The process pauses if initial health rises above zero at any step, giving positions a chance to rebound.

## NLP Liquidation Priority

The [Nado Liquidity Provider (NLP)](/nlp) vault actively participates in liquidations to maintain platform stability and generate yield for liquidity providers. When a liquidatable position is detected, the NLP may front-run external liquidators under certain conditions:

* **Perpetual Positions Only**: The NLP will attempt to front-run liquidations for perpetual futures positions. The NLP does not liquidate spot positions or spread positions.
* **Risk-Based Throttling**: When the NLP vault's risk exposure exceeds a predefined threshold, it will stop front-running liquidations, allowing external liquidators to process them instead. This ensures the NLP maintains healthy risk parameters.

Maintenance Margin Usage is an indicator of when liquidation begins. It indicates the percentage of your maintenance margin that is consumed by open positions. It provides a real-time gauge of how close your account is to liquidation thresholds.

* **Low Risk**: 0 – 40%
* **Medium Risk**: 40 – 70%
* **High Risk**: 70 – 90%
* **Extreme Risk**: 90 – 100%

> If Maintenance Margin Usage reaches 100%, your account is immediately eligible for liquidation. At this point, you will be unable to open new positions until some margin is freed up, either through position closures, additional deposits, or positive PnL realizations.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-7de486e5e982e7a9ce16149425516a459c06dff6%252FDocs-2.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=5172fa0c&sv=2)

## Liquidation Sequence

Liquidators specify a product and amount to target, with the system rounding down to the optimal size that brings maintenance health back to non-negative, whilst making sure the initial health is non-positive. – balancing efficiency with user protection.

The sequence of liquidation operations is as follows:

1. **Cancel Open Orders**: All pending orders in the subaccount are voided to free up resources.
2. **Liquidate Assets**: Spot balances, long spreads, and positive-PnL perpetuals are sold at a discount.
3. **Liquidate Liabilities**: Borrows and short spreads are repaid at a markup.

## Liquidation Price

The execution price is the oracle price adjusted by the maintenance weight divisor with a 0.5%/0.25% floor, incentivizing quick intervention.

**For Longs**:

Liquidation Price=oracle\_price×(1−max⁡(1−maint\_asset\_weight5,0.005))\text{Liquidation Price} = \text{oracle\\_price} \times \left(1 - \max\left(\frac{1 - \text{maint\\_asset\\_weight}}{5}, 0.005\right)\right)Liquidation Price=oracle\_price×(1−max(51−maint\_asset\_weight​,0.005))

**For Shorts**:

Liquidation Price=oracle\_price×(1+max⁡(maint\_liability\_weight−15,0.005))\text{Liquidation Price} = \text{oracle\\_price} \times \left(1 + \max\left(\frac{\text{maint\\_liability\\_weight} - 1}{5}, 0.005\right)\right)Liquidation Price=oracle\_price×(1+max(5maint\_liability\_weight−1​,0.005))

**Minimum Penalty: 0.5%** - At very high leverage, the liquidation penalty is capped at a minimum of 0.5% to ensure liquidators are always incentivized.

---

**For Spread Liquidations**:

When liquidating spread positions (offsetting spot and perp positions), the penalty calculation differs based on which side you're closing:

**Selling Spread (amount > 0)**: Selling spot, closing perp short

1. Calculate penalty: `penalty = (1 - perp_maint_asset_weight) / 10`
2. Apply minimum floor: `penalty = max(penalty, 0.0025)` (0.25%)
3. Liquidation price (discount): `spot_price × (1 - penalty)`

Liquidation Price=spot\_price×(1−max⁡(1−perp\_maint\_asset\_weight10,0.0025))\text{Liquidation Price} = \text{spot\\_price} \times \left(1 - \max\left(\frac{1 - \text{perp\\_maint\\_asset\\_weight}}{10}, 0.0025\right)\right)Liquidation Price=spot\_price×(1−max(101−perp\_maint\_asset\_weight​,0.0025))

**Buying Spread (amount < 0)**: Buying spot to cover short, closing perp long

1. Calculate penalty: `penalty = (spot_maint_liability_weight - 1) / 10`
2. Apply minimum floor: `penalty = max(penalty, 0.0025)` (0.25%)
3. Liquidation price (markup): `spot_price × (1 + penalty)`

Liquidation Price=spot\_price×(1+max⁡(spot\_maint\_liability\_weight−110,0.0025))\text{Liquidation Price} = \text{spot\\_price} \times \left(1 + \max\left(\frac{\text{spot\\_maint\\_liability\\_weight} - 1}{10}, 0.0025\right)\right)Liquidation Price=spot\_price×(1+max(10spot\_maint\_liability\_weight−1​,0.0025))

**Minimum Penalty: 0.25%** - Spread liquidations have a lower minimum of 0.25% (half of non-spread) since they're less risky.

---

**For non-spread liquidations, for example**, this creates a gross profit for liquidators:

Long=oracle\_price×max⁡(1−maint\_asset\_weight5,0.005)\text{Long} = \text{oracle\\_price} \times \max\left(\frac{1 - \text{maint\\_asset\\_weight}}{5}, 0.005\right)Long=oracle\_price×max(51−maint\_asset\_weight​,0.005)

Short=oracle\_price×max⁡(maint\_liability\_weight−15,0.005)\text{Short} = \text{oracle\\_price} \times \max\left(\frac{\text{maint\\_liability\\_weight} - 1}{5}, 0.005\right)Short=oracle\_price×max(5maint\_liability\_weight−1​,0.005)

But Nado allocates 50% of this to the insurance fund, so net profit is half:

Net Long=oracle\_price×max⁡(1−maint\_asset\_weight10,0.0025)\text{Net Long} = \text{oracle\\_price} \times \max\left(\frac{1 - \text{maint\\_asset\\_weight}}{10}, 0.0025\right)Net Long=oracle\_price×max(101−maint\_asset\_weight​,0.0025)

Net Short=oracle\_price×max⁡(maint\_liability\_weight−110,0.0025)\text{Net Short} = \text{oracle\\_price} \times \max\left(\frac{\text{maint\\_liability\\_weight} - 1}{10}, 0.0025\right)Net Short=oracle\_price×max(10maint\_liability\_weight−1​,0.0025)

---

## Outcomes & Insolvency

Liquidation yields two possibilities per product:

1. **Positive Outcome**: The subaccount receives USDT0 (e.g., from selling positive spot balances or closing profitable perps).
2. **Negative Outcome**: The subaccount spends USDT0 (e.g., repaying borrows or covering losing perps).

If all viable liquidations are negative and would deplete USDT0 below zero, the account becomes insolvent – positions are closed at a net loss, generating bad debt. At this point, further safeguards are activated.

**Example**

Suppose you have 31,000 USDT0 and short 10 ETH spot ($30,000 at $3000 / ETH) now. ETH spot maintenance\_liability\_weight is 1.05, then:

Maintenance Health:31,000−30,000×1.05=−500 USDT0\text{Maintenance Health}: 31{,}000 - 30{,}000 \times 1.05 = -500 \, \text{USDT0}Maintenance Health:31,000−30,000×1.05=−500USDT0

A liquidator targets 5 ETH short:

* **Oracle Price** = $3,000 / ETH
* **Liquidation Price** = $3,000 \* (1 + (1.05 - 1) / 5) = $3,030 / ETH
* **Profit** = $30 / ETH
* **Liquidation Fee** (sent to insurance fund) = $15 / ETH

Then:

* You lose 5 ETH short and the loss equals 3,030 \* 5 = 15,150 USDT0.
* You now have a 5 ETH short and 15,850 USDT0 and the maintenance health increases to 100 USDT0. You cannot be liquidated for now.
* The liquidator gets 5 ETH short with (3030 - 15) \* 5 = 15,075 USDT0 (in fact, it will be 15,074 considering $1 fee by Nado, not insurance).
* Liquidation fee (15 \* 5 = 75 USDT0) is distributed to the insurance fund.

**Note**: The liquidator cannot liquidate too much (e.g., 10 ETH short) if it will let the initial health > 0.

## High-Leverage Example (Minimum Penalty Floor)

At very high leverage, the minimum penalty floor becomes active. Suppose you have a 50x leveraged long SOL position:

* **Position**: Long 1,000 SOL-PERP at $100/SOL (entered when SOL was $100)
* **Current Price**: $100/SOL (oracle price)
* **Maintenance Asset Weight**: 0.99 (for 50x leverage)
* **Your USDT0**: $500 (maintenance health is now negative)

The natural penalty would be:

Natural Penalty=1−0.995=0.015=0.002=0.2%\text{Natural Penalty} = \frac{1 - 0.99}{5} = \frac{0.01}{5} = 0.002 = 0.2\%Natural Penalty=51−0.99​=50.01​=0.002=0.2%

But since 0.2% < 0.5% minimum, the **0.5% floor applies**:

* **Liquidation Price** = $100 × (1 - 0.005) = $99.50/SOL
* **Gross Profit** = $100 × 0.005 = $0.50/SOL
* **Liquidation Fee** (to insurance) = $0.25/SOL
* **Liquidator Net Profit** = $0.25/SOL

If a liquidator takes 500 SOL:

* You lose 500 SOL-PERP and pay 500 × $99.50 = $49,750 USDT0
* You now have 500 SOL-PERP long and $50,250 USDT0 (maintenance health restored)
* Liquidator gets 500 SOL-PERP for 500 × $99.50 = $49,750 USDT0 (market value: $50,000)
* Insurance fund receives $125 USDT0

Without the 0.5% floor, liquidators would only earn 0.2% ($100 total), making them unwilling to act. The floor ensures liquidations remain profitable even at extreme leverage.

---

## Insurance Fund & Socialization

The insurance fund is Nado's buffer against insolvency – a segregated USDT0 pool that covers bad debt by paying liquidators to absorb underwater positions.

Seeded by the Nado team, it's replenished with 50% of liquidation profits, growing with platform activity.

If depleted, losses socialize in tiers to spread impact equitably:

1. **Perpetual Socialization**: Pro-rata deductions from other positions in the same market (e.g., all ETH-PERP holders share a 0.1% collateral trim).
2. **USDT0 Depositor Socialization**: Remaining shortfalls hit USDT0 balances across the platform.

This multi-layer approach – insurance first, then targeted, then broad – ensures platform resilience and solvency amid extreme volatility.

## Example — Insolvency Resolution

Post-liquidation, your subaccount owes $1,000 bad debt on a liquidated SOL-PERP.

* The insurance fund ($1M total) injects $1,000, paying a liquidator $1,050 to take the position (5% incentive).
* The insurance fund’s capital is reduced accordingly to cover the liquidator payment.
* If the insurance fund is depleted, SOL-PERP holders each lose 0.01% of margin (pro-rata).

Users are encouraged to monitor health proactively via the Nado app. With these liquidation mechanisms, Nado turns potential tempests into manageable swells – protecting capital while enabling confident trades.

---


Last updated 16 days ago

---


# Maintenance Windows

Source: https://docs.nado.xyz/maintenance-windows

Nado features two fixed maintenance windows each week to deploy technical updates, new features, product enhancements, UX/UI improvements, and optimizations.

These windows are intentionally scheduled during periods of lowest market activity to minimize disruption. Users should expect limited app functionality for up to 10 minutes once a window begins – comparable to downtime on major CEXs and DEXs. In rare cases requiring more extensive updates, downtime may exceed 10 minutes; advance notice will be posted approximately one hour prior in the Official Nado Comms Channels.

> User assets remain fully secured on-chain and are never at risk during a maintenance window.

Not every scheduled window is utilized. A Status Bar in the bottom-left corner of the Nado app interface clearly indicates system state:

* **Operational** = All systems are operational.
* **Maintenance** = An update is ongoing (degraded app performance is possible).

For detailed real-time monitoring, visit the dedicated [Nado Status Page](https://nado-xyz.betteruptime.com/), which displays current and historical uptime for both the application and API, upcoming scheduled maintenance details, and monthly-categorized incident reports since mainnet launch.

The fixed maintenance schedule (in UTC) is as follows:

* **Window #1**: Mondays at 14:00 UTC
* **Window #2**: Thursdays at 14:00 UTC

Automated traders using the Nado API / SDK should account for these windows, as connections may experience brief downtime. For questions or coordination, reach out via support ticket or directly to admins in the Official Nado API Client Channel.

---


Last updated 1 month ago

---


# Margin Types

Source: https://docs.nado.xyz/margin-types

Nado offers two margin modes to suit your trading style:

1. **Unified Margin**
2. **Isolated Margin**

Unified margin is a form of cross-margin that treats your entire account – deposits, positions, and unrealized profits / losses – as a single, interconnected pool of collateral across spot, perpetuals, and money markets (e.g., spot borrowing). A single collateral pool allows assets to automatically offset risks via a single health score that dynamically adjusts to your portfolio’s collateral – computed via the on-chain risk engine.

> Unified margin is ideal for diversified portfolios where you want to maximize capital efficiency without intensive manual adjustments.

Isolated margin, on the other hand, assigns a fixed amount of collateral to one specific perpetual position, keeping it separate from the rest of your account – like placing a single bet in a sealed compartment to protect the rest of your funds. Isolated margin is widely popular for confining collateral risks to a single perpetual position, where margin maintenance and liquidation risks only impact that specific open position.

> Isolated margin is useful for high-volatility trades, such as using higher leverage or opening positions in more price-volatile altcoin perpetuals. If you prefer strict risk limits per position, then isolated margin enables you to ensure that one trade’s outcome doesn’t affect others.

Both modes are available on Nado using the same account – with unified cross-margin as the default margin type. Users can switch seamlessly between unified cross-margin and isolated margin via the Nado order panel to adapt their strategy as markets shift.

---

## Unified Margin

Unified margin is a form of cross-margin that consolidates all your account balances and open positions into one margin pool, enabling real-time margin offsets that reduce overall margin requirements and enhance capital efficiency.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-0353640755334357d898090465457ebf04e07a5e%252FDocs.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=a72058fb&sv=2)

## How it Works

* **Shared Collateral Pool**: Every asset in your Nado account (e.g., USDT0 deposits, kBTC holdings, etc) and all open positions contribute to a unified health score, calculated by Nado's on-chain risk engine. This score reflects your total available collateral and maintenance margin levels before hitting margin limits and liquidation thresholds.
* **Automatic Netting & Rebalancing**: Positive PnL from one position can offset losses in another, dynamically adjusting margin needs. Maintenance margin (the buffer against liquidation) and initial margin (required to open positions) are computed holistically, often resulting in lower margin requirements than isolated margin trades.
* **Risk Tiers for Visibility**: Monitor your portfolio through intuitive risk levels, updated in real-time for proactive management.

## Benefits & Examples

The core advantage of unified cross-margin is maximizing capital efficiency.

Your portfolio operates as an interconnected system, where balanced exposures minimize tied-up funds and amplify buying power. It's suited for strategies like basis trades or multi-asset hedges, eliminating the need for constant rebalancing.

**Example**: Consider the following basis trade scenario:

* You deposit 10,000 USDT0 and open a long wETH spot position worth $20,000 (2x leverage).
* Simultaneously, you short an ETH perpetuals contract for $20,000 notional (also 2x leverage).
* With unified margin and custom logic for handling spread trades, offsetting exposure nets out – meaning your effective margin requirement may drop notably compared to an isolated margin ETH perpetual of the same size. The risk engine recognizes the hedge and adjusts the margin levels accordingly.
* More specifically, if the ETH price dips 5%, the spot loss is balanced by the perpetual’s profits, maintaining health above 80% without requiring additional collateral deposits.

As a result, unified margin not only saves capital but also automates risk adjustments during market swings, letting you scale your trades without constantly making manual changes.

## Leverage & Margin Requirements Under Unified Margin

When trading with cross-margin, the leverage selector in the order panel does not determine how much margin is reserved for your position. Instead, it acts purely as a front-end sizing control that limits your maximum notional size per order.

Under the hood, margin requirements are calculated using a market’s risk parameters. Because unified margin exposes your entire account as a shared collateral pool, Nado does not lock a specific amount of margin based on the leverage selected.

The margin shown is a visual metric to gauge the amount of margin utilized by a position.

What this means in practice:

* Selecting 5x leverage does not assign 20% of notional as margin for an order.
* Initial margin is calculated using the market’s risk weights.
* Your entire account equity is available to support the new position, rather than a fixed, isolated amount.

As a result, traders may see cases such as:

* A $4,000 notional position
* Max leverage: 20x
* Selected leverage: 5x
* Actual margin used: ~$200, because the risk engine is applying the market’s leverage parameters, not the front-end slider.

Unified margin is designed to maximize capital efficiency, allowing Nado’s risk engine to evaluate your account holistically rather than locking collateral per position. The leverage slider remains a tool for controlling order size, not a determinant of allocated margin.

---

## Isolated Margin

Isolated margin dedicates a precise amount of collateral to a single perpetual position, ensuring its risks stay contained.

The isolated margin type is available exclusively for perpetuals, with a cap of one isolated position per perpetual market. It complements unified margin by allowing users to express their trading strategies with a sharper focus, removing the position’s margin impact from their Nado account’s broader portfolio of assets.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-d8eacb0184124765acc80b84e2d508d3b42cf7ed%252FDocs%2520%282%29.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=712c7c19&sv=2)

## How it Works

* **Dedicated Allocation**: When opening a position, select isolated mode and specify the desired leverage amount. This amount alone determines the position's health and leverage (up to 20x), independent of your main Nado account.
* **Independent Risk Checks**: Each isolated position has its own liquidation threshold – no shared PnL or offsets. If your position breaches maintenance margin (typically 50% of initial), it's liquidated solely on its collateral, leaving the rest of your account intact and unaffected.
* **Integration with Cross-Margin**: Funds on Nado transfer smoothly from your unified pool to fund a new isolated margin position. Notably, traders can still utilize both unified and isolated margin in the same market (e.g., cross-margin spot wETH alongside isolated ETH perps) per their trading strategy or risk preferences.

## Benefits & Examples

Isolated margin excels in targeted risk control, limiting downside to predefined amounts while enabling bolder leverage on specific markets. It promotes precise position sizing and eases oversight of outliers, functioning as a safeguard for exploratory trades amid broader stability.

Unlike unified margin’s shared collateral pool, isolated acts as a firewall. It’s ideal for testing new strategies or riding short-term swings without portfolio-wide risk impacts.

**Example**: Consider the following SOL perpetual scenario.

* Bob starts with 10,000 USDT0 in spot assets in his account.
* Bob allocates 2,000 USDT0 to a 10x leveraged SOL-PERP long amid a volatile market – a $20,000 notional position.
* Bob’s main Nado account now holds 8,000 USDT0 in spot assets, while the 2,000 USDT0 functions as the collateral for his SOL-PERP position.
* If the SOL price drops by 10% and triggers liquidation of Bob’s SOL-PERP, Bob only loses the 2,000 USDT0 collateral.
* The remaining 8,000 USDT0 spot holdings in Bob’s main account remain untouched, preserving 80% of his capital and limiting the risk exclusively to his SOL-PERP.

Nado users can also adjust their open isolated margin positions by simply adding / removing margin via the trading terminal (with optional borrowing from money markets). This recalculates health instantly without affecting other positions.

Switch between modes anytime in the order panel, or manage isolated trades directly from the positions table. With clear health indicators and automated calculations, Nado equips you to trade with confidence, turning market flux into focused opportunity.

---


Last updated 25 days ago

---


# Market Parameters

Source: https://docs.nado.xyz/market-parameters

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252FuFCIZ9xA4JORDLRki3QH%252FTable%2520interactive%2520component%2520%282%29.png%3Falt%3Dmedia%26token%3D0c9705d6-1b6e-4a73-8115-298004080358&width=768&dpr=4&quality=100&sign=18c9d29&sv=2)


Last updated 11 days ago

---


# Mission

Source: https://docs.nado.xyz/mission

We saw a market fractured by fragmentation and compromise – slow executions, thin liquidity, and platforms built for survival, not supremacy. Born from the battle-tested fires of Kraken and the Ink Foundation, Nado rejects the chaos. We engineer the world's most performant DEX: a powerful orderbook fusing blistering speeds with perpetual and spot markets – anchored by cross-margin efficiency.

Our mission cuts deeper. We hand traders the reins to tame DeFi's turbulence – funneling raw volatility into your clear signal, your edge.

Trust your instincts. Execute with unmatched speed, depth, and precision. Don't just endure volatility’s chaos, make it yours to command.

Nado. Welcome to the perfect storm.

---


Last updated 1 month ago

---


# NLP

Source: https://docs.nado.xyz/nlp

In decentralized finance, perpetual futures markets blend high potential with intense volatility.

Liquidity in altcoin pairs often lags behind major markets, fostering gaps in depth and efficiency that complicate trading. The Nado Liquidity Provider (NLP) aims to counter this within Nado's CLOB DEX, via a vault architecture that bolsters infrastructure and delivers yields to liquidity providers (LPs).

> NLP transforms USDT0 deposits from idle capital into active liquidity across the exchange.

Seamlessly embedded in Nado's orderbook, it directs LP capital toward bespoke vault strategies deployed across Nado’s perpetual markets. LPs accrue yield while narrowing spreads, enhancing order fill execution, and minimizing slippage.

For retail LPs or traders, the NLP offers a path to competitive APYs and better liquidity, turning altcoin volatility into shared platform strength.

---

## How NLP Works: From Deposit to Dynamic Yield

Under the hood, NLP operates as a federation of sub-vaults, each tied to a dedicated strategy and funding is allocated across these pools by predefined weights — overseen by weighted subaccounts

LP capital is distributed proportionally, with automated rebalancing – calculated off-contract to minimize proportional errors – occurring seamlessly during deposits and withdrawals to preserve equilibrium.

Yields are sourced from unrealized PnL from the vault strategy deployed in perpetual markets, harnessed through liquidations and market-making while prioritizing capital safeguards and swift redemptions over leveraged exposures.

> Users share equally in all PnLs, creating a diversified buffer that tempers volatility into more predictable passive income. As outcomes compound, user LP shares generate yield on USDT0 deposits proportional to the total share of LP capital deposited into the vault.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-d798fc0daf95ca378eaf7c5e8007e825b82e6e30%252FDocs-5.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=a22da6e5&sv=2)

To make the LP token pricing mechanism more accessible, consider it like determining the value of your portion of a shared pie where the size of each slice adjusts based on the pie’s total worth. Here's a step-by-step explanation:

* **Calculate Total Assets**: Begin with the total USDT0 deposited into the vault – this reflects the collective USDT0 contributions from all LPs.
* **Incorporate Unrealized PnLs**: Next, factor in the current estimated profits or losses from open perpetual positions. These "unrealized" values are calculated using up-to-date prices from oracle feeds, providing a realistic snapshot of ongoing trades without waiting for them to close.
* **Divide by Outstanding LP Tokens**: Finally, divide the combined total (assets plus unrealized PnLs) by the number of LP tokens in circulation. Each LP token reflects a proportional claim on this full value, much like dividing a pie amongst your friends.

This approach ensures transparency and equity. When you withdraw, you receive an amount based on the vault's value at that exact moment. It's a real-time reflection of performance.Withdrawals redeem at the current price, though a 4-day lock post-mint encourages long-term alignment, with burns available only after this period.

A modest withdrawal fee guards against oracle latency: 1 USDT0 for sequencer costs, plus the greater of 1 USDT0 or 10 bps of the amount – redistributed to remaining providers as a shared incentive for long-term alignment.

---

## Key Features: What Powers the NLP

The NLP fuses institutional-grade strategies with seamless retail access, unlocking yields tailored to Nado's altcoin prowess. These innovations help amplify DEX liquidity, sparking a self-reinforcing loop where every deposit fuels better trades across the platform.

## Accessible, Diversified Strategies

Sub-vaults are capable of deploying multiple approaches in parallel. Currently, only one vault is live for the Private Alpha — primarily utilizing a liquidations and maker spread strategy. All profit and loss (PnL) outcomes are aggregated and distributed proportionally across the vault, thereby promoting effective risk diversification and contributing to the overall stabilization of yields.

At the core of the system, strategies are calibrated to target perpetual futures contracts. By harnessing oracles for price feeds, the strategies can mitigate custody-related vulnerabilities – such as those arising from asset transfers or off-chain exposures.

Under the hood, a robust architecture of dedicated subaccounts function as specialized pools governed by weighted ownership structures.

Consider, for instance, allocations to institutional-grade traders specializing in market-making: these entities are empowered to execute orders on behalf of the pool, leveraging their expertise to generate alpha in real-time.

All ensuing profits – net of operational nuances – are then channeled back to LPs, accruing proportionally to each liquidity provider's share and fostering a collaborative ecosystem where individual contributions amplify collective gains.

> It is important to note that NLP vault yields are inherently variable, influenced by prevailing market conditions and the performance dynamics of the underlying strategies.

Moving forward, additional vaults with differing strategies can be added to the NLP to boost yield and provide more depth for altcoin perpetuals on the Nado orderbook.

## USDT0-Only Simplicity

Deposits and redemptions are streamlined exclusively through USDT0, eliminating the complexities and distractions inherent in managing multiple assets. This singular focus creates a clean, predictable entry point for liquidity providers, allowing them to engage without the encumbrance of fragmented holdings or conversion overheads.

> **Important**: The value of the NLP position is capped at a maximum of 20,000 USDT0 per Nado trading account during the Private Alpha.
>
> For instance, if a user deposits 19,000 USDT0 and the position grows to 20,000 USDT0, they would reach the cap. Any changes to these caps will be communicated ahead of time and updated on the Nado app interface either during the Private Alpha or later.

## Calibrated Risk Controls

Real-time oracle updates keep share prices sharp, with fees used as a latency buffer. Health gates enforce safe leverage, pausing withdrawals if needed for deleveraging. The vault's design ensures it cannot be liquidated – even in unhealthy states – allowing compounding to endure volatility without exposure to forced closures.

NLP never liquidates spot assets, but NLP pool owners can trade spot assets, minimizing risks associated with custody and asset transfers. The NLP’s design, where the NLP cannot be liquidated even in an unhealthy state, allows compounding to better endure volatility.

The 4-day lock post-mint fosters better longer-term commitment, balancing liquidity with stability.

## Altcoin Liquidity Amplifier

The NLP strategically targets long-tail perpetual futures contracts – less active trading pairs characterized by wider bid-ask spreads and shallower order book depths, where traditional liquidity provision often falls short.

The protocol's core advantage emerges by allocating capital to these altcoin pairs, where vault strategies actively contribute to improved liquidity and execution in those markets – helping to narrow spreads under more sustained book depth.

Additionally, the NLP helps accelerate new listings, bridging the gap between emerging assets and eager traders. Each USDT0 deposit into the vault acts as a catalyst, transforming idle capital into a dynamic way to enhance trading liquidity while returning yield.

---

In the perpetual landscape, where volatility unfolds with unyielding intensity, liquidity emerges as the defining differentiator – facilitating seamless order matching, reducing slippage, and enabling the precision required for sustained edge extraction.

> NLP democratizes access to these vault-caliber tactics for retail liquidity providers, delivering real yield for passive capital allocation.

A core tenet of NLP is its insulation from the perils of isolated leverage exposure, such as unanticipated liquidations that can erode individual positions. The vault's fortified architecture – bolstered by diversified sub-strategies – serves as a resilient intermediary.

Potential disruptions, whether from oracle discrepancies or transient strategy under-performance, are attenuated through layered fee structures and gating mechanisms that distribute impacts across the LP base, ensuring no single event overwhelms the system.

In summary, the NLP turns idle USDT0 deposits into more active capital on the platform, compounding yields and improving platform liquidity via market depth on less liquid altcoin pairs.

---


Last updated 17 days ago

---


# Onboarding Tutorial

Source: https://docs.nado.xyz/onboarding-tutorial

This comprehensive guide walks you through the essential steps to set up your wallet, bridge assets to the Ink L2 network, and begin trading on Nado. Whether you're new to decentralized trading or transitioning from another platform, follow these instructions sequentially for a seamless experience.

> Nado operates on the Ink L2 network, a high-performance EVM-compatible layer two network optimized for DeFi applications and powered by the Optimism Superchain.

By the end of this tutorial, you'll have a funded wallet connected to Ink and ready to execute your first trade on Nado. Let's dive in.

---

## Step 1. Choose Your Wallet

Nado and the Ink network support a variety of popular wallets, ensuring compatibility with both desktop and mobile users. Select the one that best aligns with your needs for security, usability, and features:

* **MetaMask** (Browser + Mobile): The most widely used wallet for EVM-compatible chains, ideal for beginners with its intuitive interface.
* **Rabby Wallet** (Browser): Tailored for active traders, offering automatic network switching, enhanced transaction simulations, and improved security previews.
* **WalletConnect** (All Compatible Wallets): Connect mobile-first options such as:

  + Trust Wallet
  + OKX Wallet
  + Rainbow Wallet
  + Crypto.com DeFi Wallet
  + Ledger Live (via WalletConnect for hardware security)

If you're handling significant balances, prioritize hardware-integrated options like Ledger for added protection.

---

## Step 2. Add the INK Network to Your Wallet

To interact with Nado on the Ink network, configure your wallet to recognize Ink as a custom chain. Use one of the following methods:

## **Method 1**: One-Click Addition

1. Visit [ChainList](https://chainlist.org/).
2. Search for "Ink."
3. Select the Ink network and click “Add to MetaMask” (or your chosen wallet). This automatically populates the required details.

## **Method 2**: Manual Configuration

If the one-click method isn't available, add the network manually in your wallet settings:

* **Network Name**: INK
* **RPC URL**:<https://rpc-gel.inkonchain.com>
* **Chain ID**: 57073
* **Currency Symbol**: ETH
* **Block Explorer URL**:<https://explorer.inkonchain.com>

Once added, switch your wallet to the Ink network via the network drop-down menu. Confirm the connection by checking your wallet balance (it should display ETH as the native token).

---

## Step 3. Acquire ETH on INK for Gas Fees

All transactions on the Ink network require a small amount of ETH to cover gas fees. Even minimal trades or deposits will incur these costs, so start with at least 0.01 – 0.05 ETH to cover initial activities both on the Ink L2 network and when interacting with Nado.

## Option A: Receive ETH Directly on Ink

* Withdraw ETH from a centralized exchange (CEX), bridge protocol, different wallet you control, or a transfer from a fiat → crypto onboarding platform you prefer.
* Use your Ink wallet address as the recipient. Ensure the sending address specifies the Ink network during withdrawal to avoid sending to the wrong chain.

**Pro Tip**: Kraken offers zero-fee withdrawals to the Ink network, making it a cost-effective choice.

## Option B: Bridge ETH from Another Chain

If you don't have ETH on Ink, bridge it from a supported source chain (e.g., Ethereum, Arbitrum, Polygon, etc.) using a supported bridge:

1. Visit an official bridge supported by Nado such as the [Superbridge](https://superbridge.app/), [Bungee](https://bungee.exchange/), or [Relay](https://relay.link/).
2. Select your source chain (e.g., Base) and set Ink as the destination.
3. Choose ETH as the asset to bridge (or swap other assets like USDT0 into ETH during the process).
4. Enter the amount, review fees, and sign the transaction in your wallet.
5. Wait for confirmation – bridging typically takes 5 – 15 minutes on average with some variance depending on network congestion. Track progress on the bridge's dashboard or the Ink block explorer.

Once bridged, your ETH will appear in your Ink wallet balance.

---

## Step 4. Add Additional Assets to Your Ink Wallet

With ETH secured for gas, fund your wallet with trading collateral such as stablecoins (e.g., USDT0 or other supported tokens. This prepares you for deposits into Nado.

* Follow the same receipt or bridging process as in Step 3: Withdraw from a wallet or CEX into your Ink address or bridge from another chain.
* Supported assets for deposits on Nado include wETH, USDT0, kBTC, wBTC, and USDC.

**Important**: Wallets do not always auto-detect custom tokens.

If an asset doesn't appear in your wallet:

1. Open your wallet and navigate to Import Token or Add Custom Token.
2. Paste the token's contract address (ensure it's the Ink-specific address, not from Ethereum or another chain – verify via the [Ink Explorer](https://explorer.inkonchain.com)).
3. Confirm the import. The token should now display with its balance.

---

## Step 5. Connect Your Wallet & Start Trading on Nado

With your Ink wallet funded, you're ready to engage with the Nado app:

1. Visit the [Nado app](https://app.nado.xyz) (bookmark this URL for security).
2. Click Connect Wallet in the top-right corner.
3. Select your wallet provider (e.g., MetaMask) and ensure it's switched to the Ink network.
4. Approve the connection request in your wallet – this grants Nado read-only access to your balances without custody.
5. Navigate to the Deposit section on the Portfolio page.
6. Select an asset (e.g., USDT0), enter the amount, and confirm the transaction. Deposits are near-instant on Nado, making your collateral available for trading immediately.

To open your first position:

* Browse the available markets in the trading terminal (e.g., ETH-PERP, BTC-PERP).
* Use the order form to place a limit or market order, leveraging Nado's unified margin system.
* Monitor your positions via the Portfolio tab.

---

## Troubleshooting Common Issues

Encountering hurdles? Use this section to resolve them quickly. If issues persist, submit a support ticket for team assistance.

* **Wallet Not Switching to INK Network?** In MetaMask or Rabby, open the network dropdown menu and manually select Ink. If unavailable, re-add the network per Step 2 above.
* **"Insufficient Gas Fee" Error?** Your wallet lacks ETH for transaction costs. To resolve the issue, bridge a small amount of ETH (as low as 0.001) from another chain or wallet to your Ink address.
* **Tokens Not Showing in Wallet?** Manually import the token using the associated contract address (Step 4). Double-check it's the Ink version via the block explorer.
* **Bridge Transaction Stuck?** Monitor the status both on the [Ink Explorer](https://explorer.inkonchain.com) and the bridge you're using using your transaction hash. If pending beyond 30 minutes, it may be due to network congestion – retry with a smaller amount or contact bridge support.

For Nado-specific issues, submit a support ticket for team assistance.

---

## Security Best Practices

Security is paramount in DeFi. Adhere to these tips to help safeguard your assets:

* **Verify URLs**: Always access Nado via bookmarked links or official sources. Beware of phishing sites mimicking the Nado domain.
* **Protect Your Seed Phrase**: Never share it with anyone, including support teams. Store it offline in a secure manner.
* **Use Hardware Wallets**: For larger wallet balances, integrate Ledger or Trezor via WalletConnect to keep private keys offline.
* **Enable Protections**: Activate MetaMask's phishing detection and transaction simulation features. Rabby Wallet excels here by previewing full transaction impacts.
* **Additional Habits**: Avoid clicking unsolicited links in emails / DMs, use two-factor authentication on linked exchanges, and regularly review connected dApps in your wallet settings.

By following this tutorial, you're equipped to trade confidently on Nado.

---


Last updated 1 month ago

---


# Bridging USDT0 to Ink

Source: https://docs.nado.xyz/onboarding-tutorial/bridging-usdt0-to-ink

USDT0 is the quote asset on Nado. The native USDT0 bridge delivers fast, low-cost transfers directly to Ink — typically completing in 2-10 minutes.

In this tutorial, learn how to transfer USDT0 from another chain to Ink and deposit directly into Nado to begin trading.

---

## Step-by-Step: Bridge USDT → USDT0 on Ink

1. Go to the official USDT0 bridge: <https://usdt0.to>
2. Click the prominent ‘Transfer’ button on the homepage.
3. Connect your wallet (MetaMask, Rabby, or WalletConnect) — ensure it’s set to the chain where your USDT currently resides (e.g., Ethereum, Arbitrum, Base, etc.).
4. Select USDT as the source token and your current chain as the source network.
5. Choose Ink (USDT0) as the destination chain.
6. Enter the amount of USDT you wish to bridge.
7. Review the summary carefully:

   * “You will receive” (exact USDT0 amount on Ink, usually 1:1 minus minimal fees).
   * Estimated bridge fee and gas on destination.
   * Expected arrival time (typically 2–10 minutes).
8. Approve the USDT spend (if first time on this chain) and confirm the transaction.
9. Wait for completion — track progress directly on the bridge interface or via the [Ink Explorer](https://explorer.inkonchain.com/) using your transaction hash.

**Note:** The bridge is non-custodial and funds arrival is automatic — no claiming step needed.

---

## After Bridging: Deposit USDT0 into Nado

1. Switch your wallet to the Ink network (if not already).
2. Visit<https://app.nado.xyz> and connect your wallet.
3. Navigate to Portfolio → Deposit (or click the Deposit button in the top-right).
4. Select USDT0 from the asset list.
5. Enter the amount (or click Max) and confirm the transaction — deposits are near-instant on Nado.
6. Once confirmed, your USDT0 balance appears in the Balances table and is immediately available for spot trading, perpetuals, lending, borrowing, or depositing into the NLP vault.

You’re now fully funded with USDT0 on Ink and ready to trade on Nado.

*For more details on the USDT0 bridge, please refer to the* [*official website*](https://usdt0.to/) *and* [*documentation*](https://docs.usdt0.to/)*.*

---

*Not financial advice. Trading digital assets involves market risk, including the risk of loss. Trading with leverage magnifies gains and losses and may lead to rapid liquidation, up to and including the loss of all collateral.*


Last updated 1 month ago

---


# Oracles

Source: https://docs.nado.xyz/oracles

Oracles are a cornerstone of Nado's market pricing, bridging off-chain market data to on-chain executions with speed, security, and risk intelligence – ensuring every trade, from spot fills to perpetual funding, reflects precise value without delay, distortion, or exploitation.

Oracle feeds update continuously via Nado's sequencer, bundling prices with actions for seamless on-chain settlement on the Ink L2 – minimizing gas and maximizing precision across Nado markets.

---

## How Nado's Oracle Model Works

Nado's oracles ingest prices from diverse, high-fidelity sources, aggregating them into tamper-resistant indexes that power critical functions like collateral valuation, liquidations, and funding rates.

> [Chaos Labs](https://chaoslabs.xyz/oracles) is the oracle provider powering Nado, delivering high-throughput and risk-aware price feeds for the Nado orderbook.

Chaos Protocol is a decentralized network of nodes that performs off-chain computations with advanced anomaly detection and outlier filtering, then pushes secure, real-time updates on-chain via EVM-compatible feeds tailored to Nado's products on Ink L2. This architecture treats data as an active risk variable, enabling circuit breakers to halt anomalous feeds and protect against manipulation – all while maintaining millisecond-level latency for seamless DEX operations.

## Key Benefits

* **Low-Latency**: Chaos’ decentralized node network delivers WebSocket-driven updates rivaling TradFi speeds, minimizing stale prices and enabling real-time mark pricing without bottlenecks.
* **Robustness & Resilience**: Built-in anomaly detection and circuit breakers proactively identify and mitigate exploits, ensuring feeds remain resilient even in volatile or adversarial markets.
* **Cost-Efficiency**: Optimized aggregation reduces update overhead, supporting broad asset coverage without inflating gas costs on Ink L2.
* **Security & Decentralization**: A growing network of validators eliminates single points of failure, with zero-knowledge proofs for reserve validations and real-time risk parameter adjustments.
* **Risk Intelligence**: Beyond pricing, feeds incorporate dynamic risk signals for automated protocol tuning, enhancing solvency and capital efficiency.

---

## Price Feeds

Nado leverages Chaos Labs' Edge Price Oracles for three core feed types, each engineered for precision, performance, and market integrity. These feeds aggregate real-time price data from multiple sources including centralized and decentralized exchanges (CEXs / DEXs), with each source assigned a reliability-based weight.

> The system calculates a weighted median price (the 50th percentile by cumulative weight) rather than a simple average, providing superior resistance to outliers and manipulation.

Chaos Labs aggregates data from a minimum of 3 sources per feed (generally 5-7 for blue-chip assets like BTC and ETH, with exceptions per feed for coverage optimization). Each feed requires a minimum quorum of 3 valid sources to compute and publish. Prices are recomputed every 500ms, ensuring low-latency freshness. Validation layers – including staleness filtering, deviation detection, and the quorum requirement – are applied to ensure data integrity.

## Source Selection & Weightings

Chaos Labs selects high-fidelity sources from top centralized exchanges (CEXs) and decentralized exchanges (DEXs), assigning reliability-based relative weights normalized to sum to 1.0. Weights are generally as follows, with exceptions applied per feed for asset-specific liquidity or coverage:

**CEXs (Primary Spot / Perp Venues):**

* **Binance**: 3/11 ≈ 0.27
* **Coinbase**: 3/11 ≈ 0.27
* **Bybit**: 2/11 ≈ 0.18
* **OKX**: 2/11 ≈ 0.18
* **Others** (Kraken, Gate, Bitget combined): 1/11 ≈ 0.09 (divided equally among them, ≈ 0.03 each)

**DEXs (Equal Weighting for DeFi Depth)**:

* **Uniswap** (and forks)
* **Curve** (and forks)
* **Hyperliquid**

DEX weights are equally distributed (e.g., 1/3 ≈ 0.33 each if all three are active for a feed), blended with CEXs in the overall aggregation. This hybrid approach ensures broad market representation while prioritizing liquid venues.

## Spot Oracle Prices

These benchmark underlying asset values from leading spot venues, using a weighted median of the last trade across USD / BUSD-paired exchanges like Binance, Coinbase, and Uniswap. Chaos Labs' anomaly algorithms filter noise in real-time, smoothing volatility for accurate collateral health scores and borrow limits.

> **Example**: For ETH / USD (a blue-chip feed using 6 sources), if Binance (weight: 0.27) quotes $3,000, Coinbase (0.27) $2,995, Bybit (0.18) $3,002, Uniswap (DEX, 0.17), and two others (combined 0.11) at $3,005 and $2,990, the weighted median at the 500ms recompute interval yields $3,000 – rejecting a spiked $3,100 outlier from a low-weight source to prevent inflated valuations, after passing staleness and deviation checks with a minimum quorum of 3 sources.

## Perpetual Prices

Tailored for Nado's perpetual markets, these feeds aggregate open interest-weighted prices from major perp venues (e.g., Binance Futures, Bybit), incorporating funding rate signals and liquidity depth. Chaos’ risk component auto-adjusts for imbalances, ensuring fair settlement without premium decay distortions.

> **Example**: In a BTC-PERP with $2B open interest skewed long (using 7 sources for this blue-chip feed), the feed blends prices from high-reliability venues like Binance ($60,000, weight: 0.27), Coinbase ($59,990, 0.27), Bybit ($60,010, 0.18), Hyperliquid (DEX, equal weight ≈ 0.08), and others (combined 0.20), yielding a weighted median index of $59,980 every 500ms – triggering a +0.02% funding rate to balance positions, validated by deviation detection and a minimum quorum of 3 sources.

## Mark Price

Nado's fair-value anchor for risk management, this weighted median-derived mid-price from orderbook depth (top 2% bids / asks) across spot and perp sources provides a manipulation-resistant reference for liquidations and PnL. Chaos Labs' circuit breakers pause updates if deviations exceed 1%, safeguarding against flash crashes.

> **Example**: In a thin ETH book with weighted median mid-price at $3,000 (from spot sources like Coinbase $2,990 weight: 0.27 and perp like Bybit $3,010 weight: 0.18, plus DEXs and others totaling 0.55; using 5 sources), versus spot $2,990, a -0.01% funding applies to your 10x long on $30,000 notional – crediting $3 hourly to offset a 1% dip's $300 loss, netting -$297 and preserving your edge in flux, after staleness filtering confirms all sources are within the 500ms recompute interval and quorum of at least 3 is met.

Nado’s oracle model turns market signals into actionable clarity – robust, rapid, and resilient.

---


Last updated 1 month ago

---


# Order Types

Source: https://docs.nado.xyz/order-types

Nado equips traders with a streamlined set of order types to navigate volatility with precision – executing at market speed or guarding edges with calculated triggers. Available across spot and perpetuals, these integrate seamlessly with the orderbook, ensuring fills that align with your strategy without excess complexity.

---

## Market Order

Executes immediately at the best available price in the orderbook. Ideal for quick entries or exits in liquid markets, where speed of execution beats exact pricing. Slippage is minimized by deep liquidity, but monitor thin books to avoid wider spreads.

> **Example**: You're bullish on ETH during a sudden rally and want to buy 1 ETH right now. Place a market buy order – it fills instantly at the current ask price of $2,500, even if it slips slightly to $2,510 in a fast-moving market.

## Limit Order

Place a buy or sell at a specific price or better, resting in the orderbook until matched. Provides control over entry / exit points – buy below current price, sell above – for strategies like scaling in or taking profits. Makers are charged lower fees, rewarding patience.

> **Example**: ETH is trading at $2,500, but you believe it'll dip before rising. Set a limit buy order for 1 ETH at $2,450 or lower. If the price drops and hits your level, it executes automatically; otherwise, it sits until conditions align, earning you maker rebates if filled.

## Stop Market Order

Triggers a market order once the asset hits a predefined stop price, converting to immediate execution. Use to limit losses (stop-loss) or chase breakouts (stop-buy). In turbulent swings, it activates swiftly via oracle feeds, protecting capital without manual intervention.

> **Example**: You hold a long position in BTC at $60,000 and want to cap losses at 5%. Set a stop market sell order at $57,000. If BTC crashes to that level, it triggers a market sell, closing your position at the next available price (say, $56,900) to stem further downside.

## Stop Limit Order

Triggers a limit order at a stop price, then executes only at your specified limit or better, combining protection with price discipline. For example, sell if price drops to $50 (stop), but only at $49.50 or higher (limit). Stop-limit orders shield against price gaps while avoiding poor fills in fast markets.

> **Example**: Short on SOL at $150, you set a stop limit buy at $160 (stop) with a limit of $162. If SOL surges to $160, it places a limit buy order at $162 or less. This locks in profits if it retraces slightly, but skips execution if it gaps wildly above $162.

## TWAP (Time-Weighted Average Price)

TWAP orders slice a large order into smaller chunks executed evenly over a set duration, averaging the price to reduce market impact. Perfect for institutional volumes or gradual accumulations – specify time window and total size, and Nado handles the cadence.

TWAP orders lower slippage in volatile sessions, preserving your edge.

> **Example**: You want to accumulate 100 ETH over the next hour without spiking the price from $2,500. Set a TWAP buy for 100 ETH over 60 minutes – it breaks into ~1.67 ETH chunks every minute, averaging your entry at $2,505 despite intra-hour swings between $2,490 and $2,520.

## Scaled Orders (Coming Soon)

Deploys a ladder of limit orders across a predefined price range, with customizable price and size distributions for staggered, gradual execution. Ideal for scaling into or out of trades amid volatility, reducing slippage and averaging fills over multiple levels. Select flat, increasing, or decreasing price distribution, paired with even split, increasing, or decreasing size distribution; anchor to current market price.

Configure with Fill-or-Kill (FoK) or Immediate-or-Cancel (IOC) behavior, or allow partial fills.

> **Example**: ETH trades at $3,000; you aim to buy 10 ETH on a potential dip without rattling the book. Set a scaled buy for 10 ETH over $2,950–$2,900 (5 even orders of 2 ETH). As price falls to $2,940, the lowest fills first at $2,900; further drops trigger the rest, netting an average entry of ~$2,920 versus a single large limit's impact.

---

To promote fair and orderly markets while protecting liquidity providers from ultra-low-latency predatory strategies, Nado applies a fixed 30 ms speed bump to all non-post-only orders (i.e., any aggressive order that would immediately cross and take liquidity from the book).

* Post-only orders (limit orders explicitly flagged as post-only) are exempt and execute with zero added latency.
* The 30 ms delay is currently uniform across all accounts and trading tiers.
* Future updates may introduce tiered speed-bump adjustments based on maker volume, NLP participation, or other contribution metrics — details will be announced in advance.

This mechanism significantly reduces adverse selection risk for passive liquidity providers without materially impacting retail or legitimate high-frequency trading strategies.

---


Last updated 1 month ago

---


# Orderbook Architecture

Source: https://docs.nado.xyz/orderbook-architecture

Nado harnesses DeFi's turbulent energy into a unified edge, fusing CEX speed with on-chain sovereignty on Ink L2. Its architecture integrates spot, perpetuals, and money markets into one risk-calibrated engine. The orderbook delivers 5–15 ms latency, MEV-proof executions, and superior capital efficiency – fueling Ink L2's premier liquidity hub.

At its core, Nado's design rests on two interlocking pillars:

1. **A protocol-level clearinghouse and risk engine housed on-chain.**
2. **A powerful off-chain sequencer – the central-limit orderbook (CLOB).**

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-c5adffcf6a28dee280bd7f573283a4a0a836ecb7%252FDocs-3.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=35fe4d0&sv=2)

Threaded through Ink's EVM-compatible rails, the two pillars form a low-latency CLOB DEX where every order converges with exacting precision.

---

## The Nado Orderbook

In decentralized exchanges (DEXs) like Nado, an orderbook is the core engine powering trades. It acts like a real-time marketplace ledger, listing all open buy orders (bids) from users willing to purchase assets at specific prices and sell orders (asks) from those ready to sell.

These orders are stacked by prices:

* **Bids**: Increase as you go lower (deeper discounts for buyers)
* **Asks**: Rise higher (premiums for sellers).

This creates a transparent view of supply and demand, enabling efficient matching of trades without intermediaries.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-860bf4d5bffc73156e2e21c42313f7d32f95ac79%252FDocs-1.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=9e2254ce&sv=2)

The depth chart visual above illustrates this dynamically:

* **Green Side** (left): Represents bids – buy orders below the current market price. Denser green bands indicate stronger buying interest and liquidity at those levels, meaning large trades can execute with minimal price slippage.
* **Red Side** (right): Shows asks – sell orders above the market price. Thicker red areas signal robust selling pressure and depth.
* **Center** (mid-price line): The equilibrium point, calculated as the average of the highest bid and lowest ask. It's where the market "balances" and most trades occur.

Mid-book depth zooms in on liquidity around this mid-price – the heart of the orderbook. It measures how much volume (in assets or value) sits just on either side, showing market resilience.

Deeper mid-book (denser dots / bands near the center) means the DEX can handle bigger trades smoothly, reducing volatility. In Nado, this depth ensures fair, slippage-resistant trades, even in volatile crypto conditions.

---

## On-Chain Clearinghouse – Calibrated Control

Every position breathes through Nado's on-chain clearinghouse – a vigilant hub that computes exposures, collaterals, and offsets in real-time. Unified cross-margin spans your portfolio where liabilities net out intuitively across spot, leverage, and borrows.

Oracle feeds pulse sub-second, triggering liquidations only when thresholds breach. Orders submitted to the sequencer are batched and submitted on-chain, starving MEV opportunities and letting your trading signal cut clean through the noise.

---

## Off-Chain Sequencer – Blistering Performance

Speed defines the strike. Nado's off-chain sequencer matches orders in 5–15 ms, then batches them for seamless on-chain settlement via Ink L2. As a central-limit orderbook (CLOB), it accepts limit orders through the intuitive front-end app or an HFT-optimized API, empowering automated, high-performance strategies.

The sequencer holds no custody: assets stay locked in on-chain smart contracts, under your sole control. It cannot censor trades or forge signatures – all settlements, from withdrawals to liquidations, demand on-chain verification.

Millisecond latencies and periodic batching render MEV extraction uneconomical, shielding executions from front-running. Like an accelerator atop on-chain rails, the sequencer delivers CEX-grade matching without custody shadows.

Nado's design channels market flux into high-velocity execution – a decisive edge for traders.

---

## Nado’s Horizon: The Perfect Convergence

Nado's technical stack unifies liquidity across products, its capital efficiency turns portfolios into weapons, and the Ink L2 offers sub-second finality.

Battle-tested on chains past amidst cycles of DeFi’s volatility, and now re-engineered for Ink's frontier, Nado delivers the tools for traders to outperform the market.

---


Last updated 1 month ago

---


# PnL Settlements

Source: https://docs.nado.xyz/pnl-settlements

PnL settlements on Nado represent the ongoing realization of profits and losses from your open positions, seamlessly converting unrealized gains or losses into actual USDT0 balances without interrupting your trading flow. At its core, PnL – profit and loss – measures the difference between a position's current market value and its opening cost, providing a real-time snapshot of performance.

On Nado, the settlement process runs continuously in the background.

As perpetuals fluctuate, positive PnL (winners) draws USDT0 from negative PnL (losers) across the platform, much like a self-balancing scale that redistributes weight to keep the entire system level. This automatic mechanism ensures your account value accurately reflects economic reality, with unsettled PnL acting as a bridge between open trades and settled cash.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-dbeb869e0b43f58b85160a19ab7754870cd53598%252FPnL%2520Settlements%2520Conceptual%2520Graphic.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=5ce9a769&sv=2)

For settlements on Nado, you'll see two key metrics:

1. **Perp PnL** (aggregate across all perpetual positions)
2. **Position PnL** (for a specific trade)

These are updated in real-time, influencing your subaccount health and available margin. No manual intervention is needed – settlements happen transparently and automatically.

---

## How PnL Settlements Work

Nado breaks lifetime PnL into two components:

* **Unsettled USDT0** (pending transfers)
* **Settled USDT0** (already realized in your balance)

Settlements occur automatically whenever losers' negative PnL can fund winners' positives, altering USDT0 balances accordingly.

![](https://docs.nado.xyz/~gitbook/image?url=https%3A%2F%2F1830223543-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252FAF0FfquYfYpmCNhlEvb9%252Fuploads%252Fgit-blob-fab5c26d6f7021ce16d57d65c5c118d547157763%252FPnL%2520Settlements%2520Conceptual%2520Graphic%282%29.png%3Falt%3Dmedia&width=768&dpr=4&quality=100&sign=af32b35b&sv=2)

This continuous flow prevents distortions, maintaining equilibrium across dynamic trading activity on the platform. More specifically:

* **Unsettled USDT0**: The portion of PnL yet to transfer between accounts. Positive values signal incoming deposits to your USDT0 balance; negative ones indicate upcoming withdrawals. It fluctuates with position volatility and hold duration – not always mirroring total PnL.
* **Settled USDT0**: PnL already incorporated into your USDT0 holdings, visible on the Portfolio Overview. Over time, as settlements accumulate, your assets and borrows adjust to reflect these changes.

Total PnL=Unsettled USDT0+Settled USDT0\text{Total PnL} = \text{Unsettled USDT0} + \text{Settled USDT0}Total PnL=Unsettled USDT0+Settled USDT0

Users can view unsettled USDT0 by checking the Unsettled column in the Balances table of the Margin Manager page on the Nado app.

> For a full audit of your settlements, navigate to **Portfolio → History → Settlements Table**, which logs every transfer with timestamps and amounts.

---

## View PnL Metrics

* **Perp PnL**: Track total unrealized across perpetuals on the Portfolio Overview or Perpetual Positions pages – essential for gauging subaccount health.
* **Position PnL**: Drill into specifics via the Perp Positions table, showing individual trade performance.

Account Value≠Assets−Borrows±Perp PnL\text{Account Value} \neq \text{Assets} - \text{Borrows} \pm \text{Perp PnL}Account Value=Assets−Borrows±Perp PnL

**Instead**:

Account Value=Assets−Borrows±Unsettled USDT0\text{Account Value} = \text{Assets} - \text{Borrows} \pm \text{Unsettled USDT0}Account Value=Assets−Borrows±Unsettled USDT0

This formula captures the dynamic nature of settlements, ensuring your displayed equity aligns with on-chain reality.

---

## Examples of Settlement in Action

Settlements aren't tied directly to a position's overall direction. Instead, they depend on market timing and peer activity. A winning trade might yield negative unsettled USDT0 if it funded others mid-hold, while a loser could show positive if it received from bigger shorts.

You open a 5 ETH long perpetual at $3,000 (notional $15,000, 5x leverage on $3,000 margin) and a 3 ETH short at $3,000.

* ETH rises to $3,100
* Your long gains $500 PnL (+3.33%)
* Your short loses $300 (-3.33%)
* Unsettled USDT0: +$200 net (winners pull from platform losers)

Over multiple hours, $150 settles into your USDT0 (from external shorts funding your long), leaving $50 unsettled. Your account value rises to $3,200 (original $3,000 + $200 net PnL), but assets now include the settled $150.

## Volatility Swing Examples

Consider the scenario if ETH dips to $2,900 mid-day.

* Your long now is - $500 PnL, short +$300.
* Unsettled flips to -$200 (your short funds platform longs).
* By evening's recovery to $3,050, it rebounds to +$150 unsettled.

> **Total Lifetime PnL**: +$250, with $100 settled – demonstrating how swings create interim negatives even on net winners.

---

## Closing Positions & Balance Adjustments

When you close a position, any remaining unsettled USDT0 settles automatically into your USDT0 balance, typically within minutes. This finalizes the trade's economics where gains boost collateral for new plays, losses deduct without surprise.

Open positions cause gradual USDT0 balance shifts as back-end settlements process – normal behavior that doesn't restrict trading. Your buying power remains tied to health, not just balances, so monitor both to better gauge your risk profile.

## Handling Negative PnL Without USDT0

If negative PnL accrues without sufficient USDT0, settlements continue but only as long as your subaccount health stays positive, Nado auto-borrows the shortfall from its embedded money markets at the prevailing rates. This keeps positions viable during drawdowns, but it adds interest.

With PnL settlements woven into Nado's fabric, your trades evolve from static bets to living strategies – precise, adaptive, and always aligned with the market's pulse.

---


Last updated 1 month ago

---


# Products

Source: https://docs.nado.xyz/products

## Spot Trading

To get started, deposit supported collaterals such as USDT0 or wETH into your unified account. These assets serve as margin, enabling up to 5x leverage on spot positions for amplified exposure without needing separate approvals.

For example, you can swap wETH for USDT0 seamlessly or open a leveraged long on wETH using your USDT0 balance.

> Capital efficiency is core to Nado's design. All deposits earn automatic native yields, compounded from platform fees and lending activity.

Unlike traditional DEXs with isolated pools, Nado unifies liquidity across products, so your spot holdings can directly offset risks in perpetuals or fund borrows from money markets. This turns spot trading into a strategic anchor – borrow to scale positions or hedge against perpetuals for balanced plays – all within one account.

---

## Perpetual Trading

Perpetual futures on Nado let you go long or short on assets like BTC, ETH, SOL, BNB, and XRP with up to 20x leverage and USDT0 settlement.

Positions are funded through your unified margin system, where the entire portfolio, including spot holdings, borrows, and other open perp positions, nets against exposures in real time. A long wETH spot position, for instance, can collateralize a short SOL perp, reducing overall risk and freeing up capital.

> Basis trading is streamlined where users can open a spot-perp pair directly to capture convergence opportunities, executing delta-neutral strategies without bridging assets or managing multiple interfaces. This simplifies complex setups, like funding rate arbitrage, into single, atomic trades.

All perps benefit from MEV protections and oracle-based pricing to prevent manipulation, ensuring fair execution amid market swings.

---

## Money Markets

Nado's money markets enable lending, borrowing, and yield accrual, intertwined with your trades for frictionless efficiency.

Lenders supply USDT0 or wETH to the pool, earning a proportional share of borrower interest as passive yield on idle assets. Borrowers access spot collateral-backed loans at dynamic rates shaped by supply and demand – rates rise with borrowing pressure to curb excess, and they fall amid liquidity abundance to spur activity.

> All loans remain overcollateralized via smart contracts: post crypto assets as collateral, held securely on-chain until repayment.

Unified margin nets exposures, so healthy spot or perp positions cover borrows without forced liquidations.

---

## Dynamic Interest Rate Model

Nado's money markets use a dynamic interest rate model to adapt rates in real time to supply and demand, promoting balanced utilization and efficient capital flow.

Every 15 minutes, the system recalculates rates based on the utilization ratio (R), defined as the proportion of borrowed assets to the total available supply in the pool:

R=borrowedborrowed+availableR = \frac{\text{borrowed}}{\text{borrowed} + \text{available}}R=borrowed+availableborrowed​

This frequent adjustment ensures rates respond swiftly to market shifts. As borrowing demand rises and R approaches 1, borrow rates climb to discourage over-leveraging, while deposit rates may dip to attract more liquidity. Conversely, when R falls (abundant supply), borrow rates decrease to incentivize uptake, and deposit rates rise to reward lenders.

For each spot product (e.g., wETH or USDT0), the model employs four tunable parameters to shape the rate curve:

* **small\_cap**: The utilization threshold below which rates remain flat at a minimal "floor" level, encouraging baseline borrowing without over-penalizing low demand.
* **large\_cap**: The high-utilization ceiling where rates plateau or escalate sharply to cap extreme borrowing.
* **floor**: The minimum borrow rate (as a percentage), setting a baseline cost even in low-demand scenarios.
* **inflection**: The pivot point where the curve bends from gradual to steeper increases, fine-tuning sensitivity around moderate utilization (typically 50 - 80%).

These parameters create a "kinked linear" curve for the optimal rate. It stays constant (at the floor) for R below small\_cap, then slopes linearly upward to the inflection point, accelerates more aggressively toward large\_cap, and caps thereafter.

This design prevents rates from stagnating at zero (discouraging idle capital) or spiking uncontrollably (avoiding liquidity crunches).

The final rates incorporate a protocol fee (interest\_fee = 0.2, or 20%) to sustain the platform:

Nado’s model fosters resilience by auto-balancing the pool, rewarding patient lenders with competitive yields, and keeping borrowing accessible without fixed hurdles. All while netting exposures across your unified account for seamless integration with spot and perpetual trades.

---

## Nado Liquidity Provider (NLP)

Nado's Liquidity Provider (NLP) transforms idle USDT0 into dynamic liquidity for perpetual markets, fueling tighter spreads and smoother executions across altcoin pairs. As a vault-based protocol embedded in the CLOB DEX, NLP channels deposits into diversified strategies that capture yields from liquidations and market-making — while netting risks in real-time for resilient capital efficiency. This complements money markets by extending passive income opportunities, turning volatility into shared platform strength without isolated leverage exposures.

At its core, NLP federates sub-vaults weighted by strategy, with automated rebalancing ensuring proportional allocations during deposits and withdrawals. Yields accrue from unrealized PnL in perpetual positions, aggregated equitably across LPs and reflected in real-time token pricing:

total assets+oracle-fed unrealized gainsoutstanding lp shares\frac{\text{total assets} + \text{oracle-fed unrealized gains}}{\text{outstanding lp shares}}outstanding lp sharestotal assets+oracle-fed unrealized gains​

Withdrawals redeem at current value after a 4-day lock to align long-term commitment, buffered by a modest fee (1 USDT0 fixed plus the greater of 1 USDT0 or 10 bps) that redistributes to holders and guards against oracle latency.

> Capped at 20,000 USDT0 per account during Private Alpha, this USDT0-only design strips away multi-asset friction, delivering clean, predictable access.

Key safeguards include health gates that pause outflows for deleveraging, non-liquidatable vaults to weather storms, and oracle-driven price fidelity — ensuring compounding endures without forced closures. Strategies target long-tail altcoins, liquidations, narrowing bid-ask gaps and accelerating new listings, while diversified subaccounts (like maker spreads) amplify depth without custody risks.

Variable yields, shaped by market dynamics, reward retail LPs with institutional-grade alpha, creating a virtuous cycle: every deposit enhances liquidity, compounding returns and precision for all traders on Nado.

---

## The Edge Converges

Nado's products form a cohesive ecosystem – fully integrated at the protocol level for seamless capital efficiency, not bolted-on features. Unified liquidity aggregates depth across spot, perpetuals, and money markets, reducing slippage and enabling reliable fills even in thin conditions.

Cross-margining treats your entire portfolio as collateral, netting exposures to free up funds and minimize over-collateralization.

On Nado, portfolios become precision tools.

---


Last updated 1 month ago

---


# Subaccounts & Health

Source: https://docs.nado.xyz/subaccounts-and-health

Subaccounts and health form the backbone of risk management on Nado, allowing traders to segment trading activities while maintaining a clear view of their overall exposure per each subaccount.

> Health is Nado's weighted measure of account stability, factoring in asset quality, volatility, and liquidity to determine your buffer against liquidations and capacity for new trades.

By default, subaccounts use unified margin for efficiency, but you can toggle isolated margin per position. Health calculations update in real-time via the on-chain risk engine, displayed as intuitive gauges in the Nado app. Green for safe sailing, red for imminent gales.

---

## Subaccounts

A subaccount is essentially a dedicated trading compartment tied to your main wallet address, enabling up to four subaccounts per address on Nado. Notably, the 4 subaccounts is the cap for users on the front-end, while API users have no limit on the number of subaccounts per wallet address.

> Subaccounts let you divide your wallet into up to four independent silos – each with its own balances, positions, and margin. Each subaccount operates autonomously.

## Subaccount Naming

When working with Nado, subaccounts follow a specific naming convention:

* `default` - This is your main/primary subaccount. Every wallet address automatically has a "default" subaccount upon first interaction.
* `default_1` - Your first additional subaccount (automatically assigned when created)
* `default_2` - Your second additional subaccount
* `default_3` - Your third additional subaccount

When you create additional subaccounts on the frontend, they are automatically assigned these sequential names (`default_1`, `default_2`, `default_3`). You can connect up to **4 subaccounts** simultaneously: `default` (main), `default_1`, `default_2`, and `default_3`.

Deposits, withdrawals, positions, and PnL remain fully isolated within it, ensuring a liquidation in one never affects the others. This structure protects diversified strategies – for instance, dedicating one subaccount to conservative spot holds and another to aggressive perpetual positions – while tying all to the same wallet address for independent oversight of each subaccount.

## Key Benefits

* **Risk Isolation**: Limit fallout from volatile or riskier trades – only the subaccount's assets are at stake during liquidation.
* **Capital Efficiency**: Within each subaccount, unified margin pools resources across spot, perps, and money markets by default – maximizing capital efficiency.
* **Compounding Ease**: Unrealized PnL from winning trades automatically bolster the subaccount's margin, fueling further opportunities without manual transfers between accounts.
* **Flexibility**: Switch between cross- and isolated-margin modes per subaccount, adapting to strategy needs without disrupting the whole. Unified cross-margin is the default account type on Nado.

To create a subaccount, navigate to the Nado app’s account menu and assign a unique label or name for that specific subaccount. Of note, users are required to make an initial transfer for each subaccount in order for it to be created. Either deposits from the user’s wallet or transfers on Nado from an existing subaccount will work.

Transfers between subaccounts are instant and incur a network fee: $1 USDT0 for standard transfers, or $0.10 USDT0 when either the sender or recipient is an isolated subaccount. For teams or advanced users, this mirrors a multi-account structure on centralized exchanges, but fully on-chain.

---

## Health

Health quantifies your subaccount's resilience, using weighted calculations to blend all assets and positions into a unified risk score. It accounts for variances like an asset's volatility (e.g., BTC vs. a stablecoin) or liquidity.

Low health signals more risk, a negative value on maintenance health triggers liquidation, and initial health gates new entries, preventing overextension.

There are two health thresholds:

1. **Maintenance Health**: Your liquidation buffer; if < 0, the subaccount risks partial or full close-out to protect the protocol.
2. **Initial Health**: Your available collateral; if < 0, you can't open new positions until collateral deposits restore it to a positive value.

In traditional terms, maintenance health approximates "USDT0 to liquidation," and initial health mirrors "free collateral." Weights adjust for these nuances. For example, stable assets get higher weights for more robust risk profiles, and volatile ones receive lower weights to reflect increased risk profiles.

## Weight Parameters

Every product (spot tokens, perps) has four weights:

* **maintenance\_asset\_weight**: Discounts / rewards assets for maintenance health.
* **maintenance\_liability\_weight**: Penalizes liabilities (e.g., shorts) for maintenance.
* **initial\_asset\_weight**: Similar for initial health, often stricter.
* **initial\_liability\_weight**: Heightens scrutiny on borrowings.

These ensure balanced contributions. For example, a high-volatility token might have a 0.8 asset weight, meaning it counts as 80% of its value toward health.

## Spot Health

Spot assets serve as Nado's primary collateral, powering trades across products. Their health is straightforward, reflecting value adjusted for stability.

Spot Health=weight×amount×price\text{Spot Health} = weight \times amount \times priceSpot Health=weight×amount×price

**Example**

You hold 5 BTC in a subaccount, with BTC at $10,000.

**Weights**:

* maintenance\_asset\_weight = 0.9
* initial\_asset\_weight = 0.8

> **Initial Health** = 5 × $10,000 × 0.8 = $40,000

> **Maintenance Health** = 5 × $10,000 × 0.9 = $45,000

With no positions, your $40,000 initial health means that much buying power for new trades. If BTC dips to $9,000, health recalculates to $36,000 initial – prompting a deposit if nearing zero.

## Perpetual Health

Perpetuals introduce leverage, so health nets current value against entry costs, capturing unrealized PnL.

**Example** (**Short Position**)

You short 5 BTC-perps at entry / entry price $10,000, current $10,000.

**Weights**:

* maintenance\_asset\_weight = 0.95
* maintenance\_liability\_weight = 1.05
* initial\_asset\_weight = 0.9
* initial\_liability\_weight = 1.1

> **Initial Health** = -5 × $10,000 × 1.1 − (-5 × $10,000) = -$5,000

> **Maintenance Health** = -5 × $10,000 × 1.05 − (-5 × $10,000) = -$2,500

This equates to 10x leverage (1 / (1 - 0.9) = 10). On other platforms, you'd see $5,000 initial margin and $2,500 maintenance – same thresholds, weighted for Nado's unified cross-margin. If BTC rises to $11,000, your short loses $5,000 PnL, dropping maintenance health further; a collateral deposit restores it.

## Leverage Calculation

Nado's weights imply maximum leverage:

For BTC's 0.9 weight, that would be 10x max – scalable per asset, ensuring volatility-matched limits.

---

## Special Cases

## Spreads

Spreads are offsetting positions on the same asset (e.g., long spot BTC, short BTC-PERP). The system recognizes spread trades as hedges, boosting your health beyond what the individual positions would provide.

> **Intuition**: If you're long spot BTC and short BTC-PERP, both positions move together. When BTC goes up, your spot gains but your perp loses. When BTC goes down, your spot loses but your perp gains. This hedging reduces risk, so the system gives you better health.

The spread health calculation follows this logic:

---

**Step 1: Calculate Basis Amount**

The basis amount is how much of your position is actually hedged (the overlapping part):

basis\_amount={min⁡(spot\_amount,−perp\_amount)if spot\_amount>0−max⁡(spot\_amount,−perp\_amount)if spot\_amount≤0\text{basis\\_amount} = \begin{cases}
\min(\text{spot\\_amount}, -\text{perp\\_amount}) & \text{if spot\\_amount} > 0 \\
-\max(\text{spot\\_amount}, -\text{perp\\_amount}) & \text{if spot\\_amount} \leq 0
\end{cases}basis\_amount={min(spot\_amount,−perp\_amount)−max(spot\_amount,−perp\_amount)​if spot\_amount>0if spot\_amount≤0​

**Example 1 - Long Spot, Short Perp**:

**Example 2 - Short Spot, Long Perp**:

---

**Step 2: Calculate Existing Weight**

This is the average penalty already applied to your positions (before considering the spread benefit):

existing\_weight=spot\_long\_weight+perp\_long\_weight2\text{existing\\_weight} = \frac{\text{spot\\_long\\_weight} + \text{perp\\_long\\_weight}}{2}existing\_weight=2spot\_long\_weight+perp\_long\_weight​

**Example** (using 20x leverage BTC):

---

**Step 3: Calculate Spread Weight**

The spread weight gives you better treatment than individual positions. It's calculated from the underlying product weights:

**Base Spread Weight**:

spread\_weight=1−1−product\_weight5\text{spread\\_weight} = 1 - \frac{1 - \text{product\\_weight}}{5}spread\_weight=1−51−product\_weight​

Where product\_weight is:

* **If spot\_amount > 0** (long spot, short perp): Use perp\_long\_weight
* **If spot\_amount ≤ 0** (short spot, long perp): Use spot\_long\_weight

**Spread Weight Caps** (to manage extreme leverage):

* **Initial health**:

spread\_weight=min⁡(spread\_weight,0.99)\text{spread\\_weight} = \min(\text{spread\\_weight}, 0.99)spread\_weight=min(spread\_weight,0.99)

* **Maintenance health**:

spread\_weight=min⁡(spread\_weight,0.994)\text{spread\\_weight} = \min(\text{spread\\_weight}, 0.994)spread\_weight=min(spread\_weight,0.994)

**Example** (using 20x leverage):

---

**Step 4: Calculate Spread Health Contribution**

spread\_health\_increase=basis\_amount×(spot\_price+perp\_price)×(spread\_weight−existing\_weight)\text{spread\\_health\\_increase} = \text{basis\\_amount} \times (\text{spot\\_price} + \text{perp\\_price}) \times (\text{spread\\_weight} - \text{existing\\_weight})spread\_health\_increase=basis\_amount×(spot\_price+perp\_price)×(spread\_weight−existing\_weight)

**Complete Example**:

**Key Takeaway**: Spreads give you better health (higher effective weight) because both legs move together, reducing risk. The system recognizes this and gives you ~5x more leverage on the hedged portion!

---

**High-Leverage Example (Spread Weight Cap)**

At very high leverage, the spread weight caps become active. Suppose you have 50x leverage positions:

Notably, health is calculated by applying the penalty first as if there is no spread, then increasing health for all the spread pairs. The existing penalty means the already applied health deduction when no spread is taken into account.

With these tools, Nado turns complexity into clarity, letting you navigate markets with measured confidence.

---


Last updated 24 days ago

---
