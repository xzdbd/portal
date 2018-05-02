import {
  LIST_ITEMS,
  DOWNLOAD_ITEM,
  COPY_LINK
} from '../constants/ActionTypes'

const initialState = []

export default function items (state = initialState, action) {
  switch (action.type) {
    case LIST_ITEMS:
      return action.items

    case DOWNLOAD_ITEM:
      return state.map(item =>
        item.id === action.id
          ? {...item, downloadCount: item.downloadCount++}
          : item
      )

    case COPY_LINK:
      return state

    default:
      return state
  }
}
