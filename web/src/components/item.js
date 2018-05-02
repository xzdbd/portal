import React from 'react'
import PropTypes from 'prop-types'
import { withStyles } from 'material-ui/styles'
import Card, { CardHeader, CardContent, CardActions } from 'material-ui/Card'
import IconButton from 'material-ui/IconButton'
import Typography from 'material-ui/Typography'
import red from 'material-ui/colors/red'
import ContentCopyIcon from '@material-ui/icons/ContentCopy'
import FileDownloadIcon from '@material-ui/icons/FileDownload'
import Tooltip from 'material-ui/Tooltip'

const styles = theme => ({
  card: {
    display: 'flex'
  },
  details: {
    display: 'flex',
    flexDirection: 'column',
    minWidth: '20%'
  },
  spacing: {
    minWidth: '60%'
  },
  actions: {
    display: 'flex',
    minWidth: '20%',
    paddingRight: 10
  },
  avatar: {
    backgroundColor: red[500]
  }
})

class RecipeReviewCard extends React.Component {
  render () {
    const { classes } = this.props

    return (
      <div>
        <Card className={classes.card}>
          <div className={classes.details}>
            <CardHeader
              title='ArcGIS_Desktop_103.zip'
              subheader='September 14, 2016'
            />
          </div>
          <div className={classes.spacing}>
            <CardContent>
              <Typography component='p' />
            </CardContent>
          </div>
          <div className={classes.actions}>
            <CardActions disableActionSpacing>
              <Tooltip id='tooltip-icon' title='复制链接'>
                <IconButton aria-label='Add to favorites'>
                  <ContentCopyIcon />
                </IconButton>
              </Tooltip>
              <Tooltip id='tooltip-icon' title='下载'>
                <IconButton aria-label='Add to favorites'>
                  <FileDownloadIcon />
                </IconButton>
              </Tooltip>
            </CardActions>
          </div>
        </Card>
      </div>
    )
  }
}

RecipeReviewCard.propTypes = {
  classes: PropTypes.object.isRequired
}

export default withStyles(styles)(RecipeReviewCard)
