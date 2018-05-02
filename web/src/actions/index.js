import * as types from '../constants/ActionTypes'

export const listItems = items => ({ type: types.LIST_ITEMS, items })
export const downloadItem = id => ({ type: types.DOWNLOAD_ITEM, id })
export const copyLink = id => ({ type: types.COPY_LINK, id })
